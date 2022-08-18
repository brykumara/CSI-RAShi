package main

import (
	c "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/brykumara/circlclone/csidh"
)

//var PrivateKeySize = 37 // Private key is a vector of length 37
var ExponentVectorLength = 74 // Exponent vector length of CSIDH 512 where B = 5
var rng = c.Reader
var rand_max = 32767

const (
	prec          = 100000
	Prime float64 = 37 * 1407181 * 51593604295295867744293584889 * 31599414504681995853008278745587832204909
)

func main() {
	var Secret = SampleSecret(Prime)
	Vec := Secret2Vec(Secret)
	fmt.Println(Vec)
}

func SampleSecret(Prime float64) float64 {
	var prime = (int64)(Prime)
	secret, err := c.Int(rng, big.NewInt(prime))
	if err != nil {
		panic(err)
	}
	Secrets := secret.Int64()
	var Secret float64 = (float64)(Secrets)
	return Secret
}

func Secret2Vec(secret float64) []float64 {
	Target := make([]float64, csidh.PrimeCount)
	for i := 0; i < csidh.PrimeCount; i++ {
		Target[i] = 0
	}
	Target[0] = secret // Initialize target vector
	B := make([]float64, csidh.PrimeCount*csidh.PrimeCount)

	// Babai Nearest Plane
	for i := csidh.PrimeCount - 1; i >= 0; i-- {
		for j := 0; j < len(B); j++ {
			B[j] = B[j] + (float64)(i*74)
		}
		TargetxB := Innerproduct(Target, B)
		Converted, _ := strconv.ParseFloat(HKZIPStrings[i], 64)
		ip1 := (TargetxB / Converted)
		ip1rounded := math.Floor(TargetxB / Converted)
		remainder := ip1 - ip1rounded
		if Compare(remainder, 0.5) > 0 {
			ip1rounded = ip1rounded + 1
		}
		r := ip1rounded
		A := make([]float64, csidh.PrimeCount*csidh.PrimeCount)
		for j := 0; j < len(B); j++ {
			A[j] = A[j] + (float64)(i*74)
		}
		Sub_Multiple(Target, A, r)
	}
	Vec := Target
	Reduce(Vec, 2, 10000)
	return Vec
}

func Compare(a, b float64) float64 {
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
