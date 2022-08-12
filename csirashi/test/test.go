package main

import (
	"math"
	m "math/rand"

	"github.com/brykumara/circlclone/csidh"
)

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

func Secret2Vec(secret uint64) []int64 {
	Target := make([]int64, csidh.PrimeCount)
	for i := 0; i < csidh.PrimeCount; i++ {
		Target[i] = 0
	}
	Target[0] = (int64)(secret) // Initialize target vector
	B := make([]int64, csidh.PrimeCount*csidh.PrimeCount)

	// Babai Nearest Plane
	for i := csidh.PrimeCount - 1; i >= 0; i-- {
		for j := 0; j < len(B); j++ {
			B[j] = B[j] + (int64)(i*74)
		}
		TargetxB := Innerproduct(Target, B)
		ip1 := (TargetxB / HKZIPStrings[i])
		ip1rounded := math.Floor(TargetxB / HKZIPStrings[i])
		remainder := ip1 - ip1rounded
		if Compare(remainder, 0.5) > 0 {
			ip1rounded = ip1rounded + 1
		}
		r := ip1rounded
		A := make([]int64, csidh.PrimeCount*csidh.PrimeCount)
		for j := 0; j < len(B); j++ {
			A[j] = A[j] + (int64)(i*74)
		}
		Sub_Multiple(Target, A, (int64)(r))
	}
	Vec := Target
	//Vec = Reduce(Vec,2,10000)
	return Vec
}

func Compare(a, b uint64) int {
	if a > b {
		return 1
	} else {
		if a == b {
			return 0
		} else {
			return -1
		}
	}
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
