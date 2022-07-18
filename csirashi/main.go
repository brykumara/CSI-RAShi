package main

import (
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/dh/csidh"
)

var rng = rand.Reader

func main() {
	var alice_secret, bob_secret [64]byte // 512 bit share
	var alice_priv, bob_priv csidh.PrivateKey
	var alice_pub, bob_pub csidh.PublicKey

	// Alice generates random secret, and then public key
	csidh.GeneratePrivateKey(&alice_priv, rng)
	csidh.GeneratePublicKey(&alice_pub, &alice_priv, rng)

	// Bob generates random secret, and then public key
	csidh.GeneratePrivateKey(&bob_priv, rng)
	csidh.GeneratePublicKey(&bob_pub, &bob_priv, rng)

	csidh.DeriveSecret(&bob_secret, &alice_pub, &bob_priv, rng)
	csidh.DeriveSecret(&alice_secret, &bob_pub, &alice_priv, rng)

	fmt.Println(csidh.Validate(&alice_pub, rng))
}
