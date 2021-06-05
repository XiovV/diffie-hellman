This is an implementation of the elliptic curve Diffie Hellman key exchange algorithm (ECDH) in Go.
It supports any curve that satisifes the `elliptic.Curve` interface.

This library provides a few methods that handle the generation of key pairs and calculation of the shared secret.

## Example
```go
package main

import (
	"crypto/elliptic"
	"fmt"
	dh "github.com/XiovV/diffie_hellman"
)

func main() {
	// creates a new ECDH instance. You can choose between P521(), P384(), P256() and P224()
	ecdh := dh.NewECDH(elliptic.P224())

	alicePublic, alicePrivate := ecdh.GenerateKeyPair()
	bobPublic, bobPrivate := ecdh.GenerateKeyPair()

	aliceShared, err := ecdh.GenerateSharedSecret(alicePrivate, bobPublic)
	if err != nil {
		panic(err)
	}

	bobShared, err := ecdh.GenerateSharedSecret(bobPrivate, alicePublic)
	if err != nil {
		panic(err)
	}

	fmt.Println("aliceShared: ", aliceShared)
	fmt.Println("bobShared: ", bobShared)
}
```

Example output:
```shell script
aliceShared:  c33d01be9e7acb9af8ce829cee481090c57a58df56cb6a0f17b4480fc7b8134b
bobShared:    c33d01be9e7acb9af8ce829cee481090c57a58df56cb6a0f17b4480fc7b8134b

```
As you can see, both aliceShared and bobShared are the same, which means that the library successfully generated shared secrets for Bob and Alice.
