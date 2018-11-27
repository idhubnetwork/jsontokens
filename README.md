# DID JSON TOKEN
DID json token is for authorization and authentication with Ethereum Elliptic Curve Digital Signature Algorithm.It's a tools for creating and managing decentralized identities, and for requesting and exchanging verified data between identities.

## EXAMPLE
DID jsontoken is used to authorization and authentication for decentralized identities.
DID JWT is a json web token with `{alg:ES256k}` in header due to signature algorithm only use Ethereum ECDSA.

## Getting started
1. Install jsontokens
```sh
go get github.com/idhubnetwork/jsontokens
```
2. Install dependencies
```sh
dep ensure -v
```

### jsontoken
* generate a new jsontoken

```go
import (
	"fmt"
	
	"github/idhubnetwork/jsontokens"
)

var privateKey = "0x6e365748e2a389106b24c241485a5308fb73548d43327c7a9fd4d972ca4cd472"

jt := NewJsonToken()

jt.Set("name", "idhub")
jt.Set("type", "test")
jt.Set("did", "idhub")

fmt.Println(jt.Get("name"))
	
err := jt.Sign(privateKey)
if err != nil {
	fmt.Println(err)
	t.Fail()
}

token, err := jt.GetToken()
if err != nil {
	fmt.Println(err)
	t.Fail()
}
	
fmt.Println(jt.ClaimJson)
fmt.Println(token)
```

* resolve a jsontoken

```go
var token = "0x7b226d7367223a2230783762323236343639363432323361323237343635373337343232326332323665363136643635323233613232363936343638373536323232326332323734373937303635323233613232373436353733373432323764222c22736967223a22307836386363386561623632323037633739313835353366396631316433643739393433303161373233303462613631616131623365643431323435356633313433303633666334663136383134343530373265346431613833643935306330373030386565623962613936373762663961656333363061343436353338616236313162227d"

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
	
fmt.Println(jt.Get("name"))
fmt.Println(jt.Get("type"))
fmt.Println(jt.ClaimJson)
fmt.Println(jt.Signature)
```

### DID JWT
Reference [jwt test](https://github.com/idhubnetwork/jsontokens/blob/master/didjwt_test.go)
