package jsontokens

import (
	"fmt"
	"testing"
)

// var privateKey = "0x6e365748e2a389106b24c241485a5308fb73548d43327c7a9fd4d972ca4cd472"
var jwtTest = "eyJhbGciOiJFUzI1NmsiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJ0ZXN0IiwibmFtZSI6ImlkaHViIiwidHlwZSI6InRlc3QifQ.QT24XXc9iDKwfclvd1K-L2A7qBeXpZtXmfRodVPsS1V5SuK7CrQDwf0AHEQnU_BNg9aS4s17ESPeEuBg_rjv_xw"

func TestGetJWT(t *testing.T) {
	jwt := NewJWT()
	jwt.Set("name", "idhub")
	jwt.Set("type", "test")
	jwt.Set("iss", "test")

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
	fmt.Println("JWT SET Success")
}

func TestJWTVerify(t *testing.T) {
	jwt := NewJWT()
	err := jwt.SetJWT(jwtTest)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = jwt.Verify()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
