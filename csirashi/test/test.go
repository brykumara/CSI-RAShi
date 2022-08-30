package main

import (
	c "crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/brykumara/circlclone/csidh"
)

var PrivateKeySize = 37
var ExponentVectorLength = 74
var rng = c.Reader
var rand_max = 32767

const (
	prec          = 100000
	Prime float64 = 37 * 1407181 * 51593604295295867744293584889 * 31599414504681995853008278745587832204909
)

func main() {
	start := time.Now()
	Secret := SampleSecret(Prime)
	Vec := Secret2Vec(Secret)
	elapsed := time.Since(start)
	fmt.Println("Random Sampling to Exponent Vector Took: ", elapsed)
	fmt.Println(Vec)
}

//go run test.go hkz.go pool.go norms.go vectors.go
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
	for i := 1; i < csidh.PrimeCount; i++ {
		Target[i] = 0
	}
	Target[0] = secret //
	C := make([]float64, csidh.PrimeCount)
	C[0] = secret
	B := make([]float64, csidh.PrimeCount*csidh.PrimeCount)
	// Babai Nearest Plane
	for i := (csidh.PrimeCount - 1); i >= 0; i-- {
		for j := 0; j < len(B); j++ {
			B[j] = B[j] + (float64)(i*74)
		}
		TargetxB := Innerproduct(Target, B)
		for j := 0; j < len(B); j++ {
			B[j] = 0
		} //
		HKZIPS, _ := strconv.ParseFloat(HKZIPStrings[i], 64) // CHANGE
		ip1 := new(big.Float).SetPrec(prec).Quo(big.NewFloat(TargetxB), big.NewFloat(HKZIPS))
		if ip1.Sign() < 0 {
			delta := -0.5
			ip1.Add(ip1, new(big.Float).SetFloat64(delta))
		}
		if ip1.Sign() < 0 {
			delta := 0.5
			ip1.Add(ip1, new(big.Float).SetFloat64(delta))
		}
		ip1int, _ := ip1.Int(nil)
		ip1floor := new(big.Float).SetInt(ip1int)
		remainder := new(big.Float).SetPrec(prec).Sub(ip1, ip1floor)
		if remainder.Cmp(big.NewFloat(0.5)) > 0 {
			ip1floor = ip1floor.Add(ip1floor, new(big.Float).SetFloat64(1))
		}
		r, _ := ip1floor.Float64()
		A := make([]float64, csidh.PrimeCount*csidh.PrimeCount)
		for j := 0; j < len(A); j++ {
			A[j] = A[j] + (float64)(i*74)
		}
		for j := 0; j < len(C); j++ {
			C[j] = C[j] - A[j]*r
		}

	}
	Vec := C
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
