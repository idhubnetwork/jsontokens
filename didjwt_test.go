package jsontokens

import (
	"fmt"
	"testing"
)

// var privateKey = "0x6e365748e2a389106b24c241485a5308fb73548d43327c7a9fd4d972ca4cd472"
var jwtTest = "eyJhbGciOiJFUzI1NmsiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiaWRodWIiLCJ0eXBlIjoidGVzdCJ9.D_mf_YzD1UhKshfuNFGYczCVkORd16ScHTJUFuZ2hfQhpdBEid9jA6H8a_c1vdw9xieh99QHCVeHcgGL1x_cQhs"

func TestGetJWT(t *testing.T) {
	jwt := NewJWT()
	jwt.Set("name", "idhub")
	jwt.Set("type", "test")

	fmt.Println(jwt)

	err := jwt.Sign(privateKey)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(jwt)
	token, err := jwt.GetJWT()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(token)
}

func TestSetJWT(t *testing.T) {
	jwt := NewJWT()
	err := jwt.SetJWT(jwtTest)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(jwt)
}
