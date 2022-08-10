package main

import (
	"math"

	"github.com/brykumara/circlclone/csidh"
)

func Innerproduct(a, b []float64) float64 {
	var InnerProduct float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		InnerProduct += a[i] * b[i]
	}
	return InnerProduct
}

func L1NormforOneVec(vec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += math.Abs(vec[i])
	}
	return Norm
}

func L1NormforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += math.Abs(firstvec[i] - secondvec[i])
	}
	return Norm
}

func L2NormforOneVec(vec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += vec[i]
	}
	Norm = math.Sqrt(Norm)
	return Norm
}

func L2NormforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += firstvec[i] - secondvec[i]
	}
	Norm = math.Sqrt(Norm)
	return Norm
}
