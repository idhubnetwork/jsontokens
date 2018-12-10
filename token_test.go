package jsontokens

import (
	"fmt"
	"testing"
	"time"
)

var token_pri_key = "0x06da20d5a2ffdeb3f5b6bca5199a28d63a6efec4b1b8b9c6c493fc532ed324ec"

const (
	did        = ""
	server_url = ""
	CRUD       = ""
	jwt_iss    = ""
	jwt_aud    = ""
	jwt_sub    = ""
	jwt_jti    = ""
)

func TestGetJsontokens(t *testing.T) {
	jt := NewJsonToken()
	jt.Set("did", did)
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
