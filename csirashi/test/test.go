package main

import (
	m "math/rand"

	"github.com/brykumara/circlclone/csidh"
)

var prec uint = 1000

//var PrivateKeySize = 37 // Private key is a vector of length 37
var ExponentVectorLength = 74 // Exponent vector length of CSIDH 512 where B = 5

func main() {

}

func CalculateP() uint64 {
	var Prime uint64 = 1
	for i := 0; i < csidh.PrimeCount; i++ {
		Prime = Prime * csidh.Primes[i]
	}
	Prime = 4*(Prime) - 1
	return Prime
}

func SampleSecret(Prime uint64) uint64 {
	var convert int = (int)(Prime)
	var secret = m.Intn(convert-0) + 0
	var Secret uint64 = (uint64)(secret)
	return Secret
}

func Secret2Vec(secret uint64) []uint64 {
	Target := make([]uint64, csidh.PrimeCount)
	for i := 0; i < csidh.PrimeCount; i++ {
		Target[i] = 0
	}
	Target[0] = secret // Initialize target vector
	// Babai Nearest Plane
	for i := csidh.PrimeCount - 1; i >= 0; i-- {

	}
	return Target
}

//func ModCn2Vec(secret float64, vec *[csidh.PrivateKeySize]int8) {
//	target := make([]float64, csidh.PrimeCount)
//	target[0] = secret
//
//	for i := csidh.PrimeCount - 1; i >= 0; i-- {
//		var ip1, floor float64
//		Inner_product(target, 5+i*74, ip1)
//
//	}
//
//}

// Calculates the inner product of 2 vectors of length 74
