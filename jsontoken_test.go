package jsontokens

import (
	"fmt"
	"testing"
)

var privateKey = "0x6e365748e2a389106b24c241485a5308fb73548d43327c7a9fd4d972ca4cd472"

var token = "0x7b226d7367223a223078376232323665363136643635323233613232363936343638373536323232326332323734373937303635323233613232373436353733373432323764222c22736967223a22307830636538343864653639313962386537643135313832383838306135373065623261376361316165306537623966376335353665323164646538633838313966323732326166356232353835366266353439643737386336306265653634656166643266666135333662386434336532343732323037306431636532326531613162227d"

// var tokenTestData = new(JsonToken)

func TestGetToken(t *testing.T) {
	fmt.Println(1)
	jt := NewJsonToken()
	fmt.Println(2)
	jt.Set("name", "idhub")
	jt.Set("type", "test")
	fmt.Println(3)
	fmt.Println(jt.Get("name"))
	err := jt.Sign(privateKey)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(jt)
	token, err := jt.GetToken()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(jt.ClaimJson)
	fmt.Println(token)
	fmt.Println("success")
}

func TestSetToken(t *testing.T) {
	jt := NewJsonToken()
	err := jt.SetToken(token)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(jt.Get("name"))
	fmt.Println(jt.Get("type"))
	fmt.Println(jt.ClaimJson)
	fmt.Println(jt.Signature)
}
