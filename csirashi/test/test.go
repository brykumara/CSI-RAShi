package main

import (
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/brykumara/circlclone/csidh"
)

const (
	Lambda  = 128
	message = 2
)

var rng = rand.Reader

func main() {

	GeneratePairs(message)

}

func GeneratePairs(message int) { // Create pairs based on (3) of the CSI-RAShi paper
	var secret csidh.PrivateKey
	csidh.GeneratePrivateKey(&secret, rng)
	Initial_set := make([]csidh.PublicKey, message)
	GroupActionSet := make([]csidh.PublicKey, message)
	for i := range Initial_set {
		var pub csidh.PublicKey
		csidh.GeneratePublicKey(&pub, &secret, rng)
		Initial_set[i] = pub
		csidh.GroupAction(&pub, &secret, rng)
		GroupActionSet[i] = pub
	}
	fmt.Println("Your initial indexed set is:", Initial_set)
	fmt.Println("Your set after applying the secret is:", GroupActionSet)
	fmt.Println("Your secret is:", secret)
}

func NIZKP(Initial_set []csidh.PublicKey, GroupActionSet []csidh.PublicKey, secret csidh.PrivateKey) {
	Challenge_set1 := make([]csidh.PublicKey, (message + Lambda))
	Secretset := make([]csidh.PrivateKey, Lambda)
	for j := 0; j < Lambda; j++ {
		var secret2 csidh.PrivateKey
		csidh.GeneratePrivateKey(&secret2, rng)
		Secretset[j] = secret2
		for i := 0; i < message; i++ {
			csidh.GroupAction(&Initial_set[i+j], &Secretset[j], rng)
			Challenge_set1[j+i] = Initial_set[i+j]
		}
	}

	HashInput := strconv.Itoa(Initial_set)
	for n := 0; n < (Lambda + message); n++ {
		input := strconv.Itoa(Challenge_set1[n])
		HashInput += input
	}

	// use SHA3 to create output of 128 bit length
	// Very confused about how we can use the secret for modular arithmetic?
}
