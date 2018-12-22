package jsontokens

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	did "github.com/idhubnetwork/jsontokens/contract"
	"github.com/idhubnetwork/jsontokens/crypto"
)

// A object to handle jsontoken.
type JsonToken struct {
	Claim     map[string]interface{} `json:"-",omitempty`
	ClaimJson string                 `json:"msg"`
	Signature string                 `json:"sig"`
}

// Init a JsonToken struct.
func NewJsonToken() *JsonToken {
	token := JsonToken{
		make(map[string]interface{}),
		"",
		"",
	}
	return &token
}

// Get retrieves the value corresponding with key from the JsonToken.
func (t *JsonToken) Get(key string) interface{} {
	if t.Claim == nil {
		return nil
	}
	return t.Claim[key]
}

// Set sets JsonToken[key] = val. It'll overwrite without warning.
func (t *JsonToken) Set(key string, val interface{}) {
	t.Claim[key] = val
}

// Del removes the value that corresponds with key from the JsonToken.
func (t *JsonToken) Del(key string) {
	delete(t.Claim, key)
}

// Has returns true if a value for the given key exists inside the JsonToken.
func (t *JsonToken) Has(key string) bool {
	_, ok := t.Claim[key]
	return ok
}

// Jsonify attribute and assign it to ClaimJson.
func (t *JsonToken) SignedMsg() error {
	if t.Claim == nil || len(t.Claim) == 0 {
		return errors.New("jsontoken no claim")
	}
	tmp, err := json.Marshal(map[string]interface{}(t.Claim))
	if err != nil {
		return err
	}

	t.ClaimJson = hexutil.Encode(tmp)

	return nil
}

// Get a integrated jsontoken.
func (t *JsonToken) GetToken() (string, error) {
	tmp, err := t.MarshalJSON()
	if err != nil {
		return "", err
	}

	return hexutil.Encode(tmp), nil
	// return string(tmp), nil
}

// Split a jsontoken to a JsonToken struct.
func (t *JsonToken) SetToken(token string) error {
	b, err := hexutil.Decode(token)
	if err != nil {
		return err
	}

	tmp := struct {
		ClaimJson string `json:"msg"`
		Signature string `json:"sig"`
	}{}

	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	t.ClaimJson = tmp.ClaimJson
	t.Signature = tmp.Signature

	b, err = hexutil.Decode(t.ClaimJson)
	if err != nil {
		return err
	}

	err = t.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	t.Set("signature", t.Signature)
	return nil
}

// MarshalJSON implements json.Marshaler for JsonToken.
func (t *JsonToken) MarshalJSON() ([]byte, error) {
	if len(t.ClaimJson) == 0 {
		return nil, errors.New("jsontoken has no hex json claim")
	}
	if !t.Has("signature") {
		return nil, errors.New("jsontoken has no signature")
	}
	return json.Marshal(*t)
}

// UnmarshalJSON implements json.Unmarshaler for JsonToken.
func (t *JsonToken) UnmarshalJSON(b []byte) error {
	if b == nil {
		return nil
	}

	b, err := DecodeEscapedNoBase64(b)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &t.Claim); err != nil {
		return err
	}
	return nil
}

// Sign attribute and assign signature to Signature.
func (t *JsonToken) Sign(privateKey string) error {
	if t.Has("signature") {
		return errors.New("jsontoken already signed")
	}
	t.Del("signature")
	err := t.SignedMsg()
	if err != nil {
		return err
	}

	signature, err := crypto.Sign(privateKey, t.ClaimJson)
	if err != nil {
		return err
	}
	t.Set("signature", signature)
	t.Signature = signature

	return nil
}

// Verify a JsonToken signature and signer did.
func (t *JsonToken) Verify() error {
	if len(t.Signature) == 0 {
		return errors.New("jsontoken has no signature")
	}
	if !t.Has("did") {
		return errors.New("jsontoken has no did")
	}

	address, ok := t.Get("did").(string)
	if !ok {
		return errors.New("did is not a hex string")
	}
	address = string(address[10:52])

	authentication, err := crypto.EcRecover(t.ClaimJson, t.Signature)
	if err != nil {
		return err
	}

	if authentication == address {
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
		return errors.New("invalid signature or did")
	}

	return nil
}
