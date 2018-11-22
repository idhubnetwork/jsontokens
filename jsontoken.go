package jsontokens

import (
	"encoding/json"
	"errors"

	"github.com/idhubnetwork/jsontokens/crypto"
)

type JsonToken map[string]interface{}

// Get retrieves the value corresponding with key from the JsonToken.
func (t JsonToken) Get(key string) interface{} {
	if t == nil {
		return nil
	}
	return t[key]
}

// Set sets JsonToken[key] = val. It'll overwrite without warning.
func (t JsonToken) Set(key string, val interface{}) {
	t[key] = val
}

// Del removes the value that corresponds with key from the JsonToken.
func (t JsonToken) Del(key string) {
	delete(t, key)
}

// Has returns true if a value for the given key exists inside the JsonToken.
func (t JsonToken) Has(key string) bool {
	_, ok := t[key]
	return ok
}

// MarshalJSON implements json.Marshaler for JsonToken.
func (t JsonToken) MarshalJSON() ([]byte, error) {
	if t == nil || len(t) == 0 {
		return nil, nil
	}
	return json.Marshal(map[string]interface{}(t))
}

// UnmarshalJSON implements json.Unmarshaler for Claims.
func (t *JsonToken) UnmarshalJSON(b []byte) error {
	if b == nil {
		return nil
	}

	b, err := DecodeEscaped(b)
	if err != nil {
		return err
	}

	// Since json.Unmarshal calls UnmarshalJSON,
	// calling json.Unmarshal on *p would be infinitely recursive
	// A temp variable is needed because &map[string]interface{}(*p) is
	// invalid Go. (Address of unaddressable object and all that...)

	tmp := map[string]interface{}(*t)
	if err = json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = JsonToken(tmp)
	return nil
}

func (t JsonToken) Sign(privateKey string) error {
	if t.Has("signature") {
		return errors.New("jsontoken already signed")
	}
	t.Del("signature")
	tmp, err := t.MarshalJSON()
	if err != nil {
		return nil
	}
	signature, err := crypto.Sign(privateKey, string(tmp))
	if err != nil {
		return nil
	}
	t.Set("signature", signature)
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
