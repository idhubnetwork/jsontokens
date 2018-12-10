package jsontokens

import (
	"fmt"
	"testing"
)

var privateKey = ""

var token = "0x7b226d7367223a2230783762323236343639363432323361323237343635373337343232326332323665363136643635323233613232363936343638373536323232326332323734373937303635323233613232373436353733373432323764222c22736967223a22307836386363386561623632323037633739313835353366396631316433643739393433303161373233303462613631616131623365643431323435356633313433303633666334663136383134343530373265346431613833643935306330373030386565623962613936373762663961656333363061343436353338616236313162227d"

func TestGetToken(t *testing.T) {
	fmt.Println(1)
	jt := NewJsonToken()
	fmt.Println(2)
	jt.Set("name", "idhub")
	jt.Set("type", "test")
	jt.Set("did", "test")
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

func TestVerifyJsonToken(t *testing.T) {
	jt := NewJsonToken()
	err := jt.SetToken(token)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = jt.Verify()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
