package jsontokens

import (
	"fmt"
	"testing"
	"time"
)

var token_pri_key = ""

const (
	didTest    = ""
	server_url = ""
	CRUD       = ""
	jwt_iss    = ""
	jwt_aud    = ""
	jwt_sub    = ""
	jwt_jti    = ""
)

func TestGetJsontokens(t *testing.T) {
	jt := NewJsonToken()
	jt.Set("did", didTest)
	jt.Set("expiration", time.Now().Unix()+int64(10000))
	jt.Set("destination", server_url)
	jt.Set("action", CRUD)
	jt.Set("jwt_iss", jwt_iss)
	jt.Set("jwt_aud", jwt_aud)
	jt.Set("jwt_sub", jwt_sub)
	jt.Set("jwt_jti", jwt_jti)
	err := jt.SignedMsg()
	if err != nil {
		t.Failed()
	}
	err = jt.Sign(token_pri_key)
	if err != nil {
		t.Failed()
	}
	fmt.Println(jt.ClaimJson)
	fmt.Println(jt.Signature)
}
