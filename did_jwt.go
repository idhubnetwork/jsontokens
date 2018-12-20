package jsontokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	did "github.com/idhubnetwork/jsontokens/contract"
	"github.com/idhubnetwork/jsontokens/crypto"
)

// json web token indicates RFC7519 JWT strictly.
const jsonwebtoken = "json web token"

// A object used to handle jwt.
//
// Claim is map filled by jwt k/v, include credential(JWT) attribute.
// Payload is base64 encoeded json string.
// Header is base64 encoeded json string, defalut {\"alg\":\"ES256k\",\"typ\":\"JWT\"}.
// Sig is base64 encoeded signature string.
type JWT struct {
	Claim   map[string]interface{}
	Payload string
	Header  string
	Sig     string
}

// Init a jwt struct.
func NewJWT() *JWT {
	jwt := JWT{
		make(map[string]interface{}),
		"",
		"",
		"",
	}
	return &jwt
}

func (t *JWT) Reset() {
	t.Claim = make(map[string]interface{})
	t.Payload = ""
	t.Header = ""
	t.Sig = ""
}

// Get retrieves the value corresponding with key from the did jwt claim.
func (t JWT) Get(key string) interface{} {
	if t.Claim == nil {
		return nil
	}
	return t.Claim[key]
}

// Set sets JWT.Claim[key] = val. It'll overwrite without warning.
func (t JWT) Set(key string, val interface{}) {
	t.Claim[key] = val
}

// Del removes the value that corresponds with key from the did jwt claim.
func (t JWT) Del(key string) {
	delete(t.Claim, key)
}

// Has returns true if a value for the given key exists inside the did jwt.
func (t JWT) Has(key string) bool {
	_, ok := t.Claim[key]
	return ok
}

// Sign jwt and assign signature to jwt.Sig and
//  set jwt.Header, jwt.Payload together.
func (t *JWT) Sign(privateKey string) error {
	if t.Claim == nil {
		return errors.New("no claim to sign")
	}
	if len(t.Sig) != 0 {
		return errors.New("jwt already signed")
	}
	if len(t.Payload) != 0 {
		return errors.New("jwt already json marshal")
	}
	header := "{\"alg\":\"ES256k\",\"typ\":\"JWT\"}"
	t.Header = string(Base64Encode([]byte(header)))
	payload, err := json.Marshal(t.Claim)
	if err != nil {
		return err
	}
	t.Payload = string(Base64Encode(payload))
	msg := t.Header + "." + t.Payload
	tmp, err := crypto.Sign_ETH(privateKey, msg)
	sig := string(Base64Encode(tmp))
	if err != nil {
		return err
	}
	t.Sig = sig
	return nil
}

// Get a base64 encoded json web token assembled by a dot
func (t *JWT) GetJWT() (string, error) {
	if len(t.Sig) == 0 {
		return "", errors.New("jwt not signed yet")
	}
	token := t.Header + "." + t.Payload + "." + t.Sig
	return token, nil
}

// Split a json web token to a JWT struct.
func (t *JWT) SetJWT(token string) error {
	// tmp := regexp.MustCompile(`[\PP]+`).FindAllString(token, -1)
	tmp := strings.Split(token, ".")
	if len(tmp) != 3 {
		return errors.New("invalid JWT")
	}
	t.Header = tmp[0]
	t.Payload = tmp[1]
	t.Sig = tmp[2]
	fmt.Println(t)
	claim, err := Base64Decode([]byte(tmp[1]))
	if err != nil {
		return err
	}
	// claim_tmp := make(map[string]interface{})
	err = json.Unmarshal(claim, &t.Claim)
	if err != nil {
		return err
	}
	// t.Claim = claim_tmp
	return nil
}

// Verify a JWT signature and iss did.
func (t *JWT) Verify() error {
	if !t.Has("iss") {
		return errors.New("jwt has no issuer")
	}

	address, ok := t.Get("iss").(string)
	if !ok {
		return errors.New("jwt issuer is not a hex string")
	}
	address = string(address[10:52])

	msg := []byte(t.Header + "." + t.Payload)
	hash := crypto.SignHash(msg)

	sig, err := Base64Decode([]byte(t.Sig))
	if err != nil {
		return err
	}

	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return errors.New("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	authentication, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return err
	}

	if strings.ToLower(authentication) == strings.ToLower(address) {
		return nil
	}

	instance, err := did.GetDid()
	if err != nil {
		return errors.New("get did instance failed")
	}

	identity := common.HexToAddress(address)
	publickKeyType := [32]byte{}
	copy(publickKeyType[:], "veriKey")
	publickKey := common.HexToAddress(authentication)

	ok, err = instance.ValidDelegate(nil, identity, publickKeyType, publickKey)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("invalid signature or issuer")
	}

	return nil
}
