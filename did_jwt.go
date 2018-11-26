package jsontokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	did "github.com/idhubnetwork/jsontokens/contract"
	"github.com/idhubnetwork/jsontokens/crypto"
)

type JWT struct {
	Claim   map[string]interface{}
	Payload string
	Header  string
	Sig     string
}

func NewJWT() *JWT {
	jwt := JWT{
		make(map[string]interface{}),
		"",
		"",
		"",
	}
	return &jwt
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

func (t *JWT) GetJWT() (string, error) {
	if len(t.Sig) == 0 {
		return "", errors.New("jwt not signed yet")
	}
	token := t.Header + "." + t.Payload + "." + t.Sig
	return token, nil
}

func (t *JWT) SetJWT(token string) error {
	tmp := regexp.MustCompile(`[\PP]+`).FindAllString(token, -1)
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

func (t *JWT) Verify() error {
	if !t.Has("iss") {
		return errors.New("jwt has no issuer")
	}

	address, ok := t.Get("iss").(string)
	if !ok {
		return errors.New("jwt issuer is not a hex string")
	}

	msg := []byte(t.Header + "." + t.Payload)
	hash := crypto.SignHash(msg)
	sig, err := Base64Decode([]byte(t.Sig))
	if err != nil {
		return err
	}

	authentication, err := crypto.Ecrecover(hash, sig)

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
