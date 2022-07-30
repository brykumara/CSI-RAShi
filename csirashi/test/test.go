package main

import (
	"crypto/rand"
	"fmt"

	"github.com/brykumara/circlclone/csidh"
)

var rng = rand.Reader

func main() {
	GeneratePairs(6)
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
	fmt.Println(Initial_set)
	fmt.Println(GroupActionSet)
}
