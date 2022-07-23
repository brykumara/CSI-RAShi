package main

import (
	"crypto/rand"
	"fmt"

	"github.com/brykumara/circlclone/csidh"
)

const (
	N       = 0x1B81B90533C6C87B // 1982068743014369403, Fix some prime N = 4 * product(Elkies primes) - 1
	parties = 5
)

var rng = rand.Reader

func main() {
	secret := csidh.RandomFp()
	fmt.Println(secret)
	GeneratePairs(3, secret)
}

func GeneratePairs(message int, secret csidh.Fp) { // Create pairs based on (3) of the CSI-RAShi paper
	E := make([]csidh.Fp, message) // E as the set of random points
	var A csidh.FpRngGen
	var V csidh.Fp
	// var Secret = csidh.
	// Ep := make([]csidh.Fp, message) // Ep as the set of random points with group action of the secret
	for i := range E {
		A.RandFp(&V, rng)
		E[i] = V
		//	Ep[i] = E[i] + Secret // Need to add each element with secret
	}
	fmt.Println(E)
	//fmt.Println(Ep)
}
