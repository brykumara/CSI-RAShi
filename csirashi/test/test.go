package main

import (
	c "crypto/rand"
	"fmt"
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
	Vec := Secret2Vec(100)
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

func Secret2Vec(secret float64) float64 {
	Target := make([]float64, csidh.PrimeCount)
	for i := 1; i < csidh.PrimeCount; i++ {
		Target[i] = 0
	}
	Target[0] = secret //
	B := make([]float64, csidh.PrimeCount*csidh.PrimeCount)
	// Babai Nearest Plane
	for i := (csidh.PrimeCount - 1); i >= 70; i-- { // 74
		for j := 0; j < len(B); j++ {
			B[j] = B[j] + (float64)(i*74)
		} //
		TargetxB := Innerproduct(Target, B) //
		for j := 0; j < len(B); j++ {
			B[j] = 0
		} // Need to reset B every time?
		HKZIPS, _ := strconv.ParseFloat(HKZIPStrings[i], 64)                                  //
		ip1 := new(big.Float).SetPrec(prec).Quo(big.NewFloat(TargetxB), big.NewFloat(HKZIPS)) //
		if ip1.Sign() < 0 {
			delta := -0.5
			ip1.Add(ip1, new(big.Float).SetFloat64(delta))
		}
		if ip1.Sign() < 0 {
			delta := 0.5
			ip1.Add(ip1, new(big.Float).SetFloat64(delta))
		}
		ip1int, _ := ip1.Int(nil)
		ip1floor := new(big.Float).SetInt(ip1int)                    //
		remainder := new(big.Float).SetPrec(prec).Sub(ip1, ip1floor) //
		if remainder.Cmp(big.NewFloat(0.5)) > 0 {
			ip1floor = ip1floor.Add(ip1floor, new(big.Float).SetFloat64(1))
		} //
		r, _ := ip1floor.Float64() // r is actually positive decreasing to 0, so we expect A*r to get smaller and smaller
		A := make([]float64, csidh.PrimeCount*csidh.PrimeCount)
		for j := 0; j < len(A); j++ { // 74*74
			A[j] = A[j] + (float64)(i*74)
			// r is locked for i, and so is i, loop over each element of A first
		}
		for j := 0; j < len(A); j++ {
			A[j] = A[j] * r
		}
		fmt.Println(A[10])
		// A*r does descend to 0, so adding them should not be a problem?
		//Sub_Multiple(Target, A, r)
		// For this, C seems to be ok. Target is just the value of C?

		//Seems to be a problem when substracting Target by A*r gets too large. Use big float?
	} //go run test.go hkz.go pool.go norms.go vectors.go
	//Vec := Target
	//Reduce(Vec, 2, 10000)
	return 0
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
