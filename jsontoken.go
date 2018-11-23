package jsontokens

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/idhubnetwork/jsontokens/crypto"
)

type JsonToken struct {
	Claim     map[string]interface{} `json:"-"`
	ClaimJson string                 `json:"msg"`
	Signature string                 `json:"sig"`
}

// Get retrieves the value corresponding with key from the JsonToken.
func (t JsonToken) Get(key string) interface{} {
	if t.Claim == nil {
		return nil
	}
	return t.Claim[key]
}

// Set sets JsonToken[key] = val. It'll overwrite without warning.
func (t JsonToken) Set(key string, val interface{}) {
	t.Claim[key] = val
}

// Del removes the value that corresponds with key from the JsonToken.
func (t JsonToken) Del(key string) {
	delete(t.Claim, key)
}

// Has returns true if a value for the given key exists inside the JsonToken.
func (t JsonToken) Has(key string) bool {
	_, ok := t.Claim[key]
	return ok
}

func (t JsonToken) SignedMsg() error {
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

func (t JsonToken) GetToken() (string, error) {
	tmp, err := t.MarshalJSON()
	if err != nil {
		return "", err
	}

	return hexutil.Encode(tmp), nil
}

func (t JsonToken) SetToken(token string) error {
	b, err := hexutil.Decode(token)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, t)
	if err != nil {
		return err
	}

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
func (t JsonToken) MarshalJSON() ([]byte, error) {
	if len(t.ClaimJson) == 0 {
		return nil, errors.New("jsontoken has no hex json claim")
	}
	if !t.Has("signature") {
		return nil, errors.New("jsontoken has no signature")
	}
	return json.Marshal(t)
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

	if err = json.Unmarshal(b, t.Claim); err != nil {
		return err
	}
	return nil
}

func (t JsonToken) Sign(privateKey string) error {
	if t.Has("signature") {
		return errors.New("jsontoken already signed")
	}
	t.Del("signature")
	err := t.SignedMsg()
	if err != nil {
		return nil
	}
	signature, err := crypto.Sign(privateKey, t.ClaimJson)
	if err != nil {
		return nil
	}
	t.Set("signature", signature)
	t.Signature = signature
	return nil
}

func (t JsonToken) Verify() error {
	if !t.Has("signature") {
		return errors.New("jsontoken has no signature")
	}
	if !t.Has("did") {
		return errors.New("jsontoken has no did")
	}
	signature, ok := t.Get("signature").(string)
	if !ok {
		return errors.New("signature is not a hex string")
	}
	tmp, err := t.MarshalJSON()
	if err != nil {
		return nil
	}
	did, err := crypto.EcRecover(signature, string(tmp))
	if did != t.Get("did") {
		return errors.New("invalid signature or did")
	}
	return nil
}
