package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	fakeJwksRSAKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	key, _ := jwk.FromRaw(fakeJwksRSAKey)
	rsaKey, _ := key.(jwk.RSAPrivateKey)
	res, _ := json.Marshal(rsaKey)
	fmt.Printf("{\"keys\":[ %s]}\n", string(res))
}
