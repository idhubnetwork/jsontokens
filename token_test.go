package jsontokens

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var token_pri_key = "0x06da20d5a2ffdeb3f5b6bca5199a28d63a6efec4b1b8b9c6c493fc532ed324ec"

var (
	count      = 1
	didTest    = "did:idhub:0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4"
	server_url = "url"
	CRUD       = "READ"
	jwt_iss    = "did:idhub:0x7EbEE9a8A3530fd1e54017C39592A5a95af99d07"
	jwt_aud    = "did:idhub:0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4"
	jwt_sub    = strconv.Itoa(count)
	jwt_jti    = strconv.Itoa(count * 58825)
)

func TestGetJsontokens(t *testing.T) {
	jt := NewJsonToken()
	jt.Set("did", didTest)
	jt.Set("expiration", time.Now().Unix()+int64(1000000))
	jt.Set("destination", server_url)
	jt.Set("action", CRUD)
	jt.Set("jwt_iss", jwt_iss)
	jt.Set("jwt_aud", jwt_aud)
	jt.Set("jwt_sub", jwt_sub)
	jt.Set("jwt_jti", jwt_jti)
	fmt.Println(jt)
	err := jt.SignedMsg()
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	err = jt.Sign(token_pri_key)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(jt.ClaimJson)
	fmt.Println(jt.Signature)
	jsontoken, err := jt.GetToken()
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println("------------------------\n------------------------\n" +
		jsontoken + "------------------------\n------------------------\n")
}
