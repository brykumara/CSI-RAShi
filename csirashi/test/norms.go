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

func L1NormSumforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += math.Abs(firstvec[i] + secondvec[i])
	}
	return Norm
}

func L1NormDiffforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += math.Abs(firstvec[i] - secondvec[i])
	}
	return Norm
}

func L2NormforOneVec(vec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += vec[i] * vec[i]
	}
	return Norm
}

func L2NormSumforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += (firstvec[i] + secondvec[i]) * (firstvec[i] + secondvec[i])
	}
	return Norm
}

func L2NormDiffforTwoVec(firstvec, secondvec []float64) float64 {
	var Norm float64 = 0
	for i := 0; i < csidh.PrimeCount; i++ {
		Norm += (firstvec[i] - secondvec[i]) * (firstvec[i] - secondvec[i])
	}
	return Norm
}

func AddVec(firstvec, secondvec []float64) []float64 {
	for i := 0; i < csidh.PrimeCount; i++ {
		firstvec[i] += secondvec[i]
	}
	return firstvec
}

func SubVec(firstvec, secondvec []float64) {
	for i := 0; i < csidh.PrimeCount; i++ {
		firstvec[i] -= secondvec[i]
	}
}
