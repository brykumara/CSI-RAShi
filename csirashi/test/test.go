package main

import (
	"crypto/rand"
	"fmt"

	"github.com/brykumara/circlclone/csidh"
)

var rng = rand.Reader

func main() {
	var alice_priv csidh.PrivateKey
	var alice_pub csidh.PublicKey

	// Alice generates random secret, and then public key
	csidh.GeneratePrivateKey(&alice_priv, rng)
	csidh.GroupAction(&alice_pub, &alice_priv, rng)

	fmt.Println(&alice_pub)
}
