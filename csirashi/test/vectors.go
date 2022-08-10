package main

import (
	"github.com/brykumara/circlclone/csidh"
)

func Sub_Multiple(target, vec []float64, mul float64) []float64 {
	for i := 0; i < csidh.PrimeCount; i++ {
		vec[i] = vec[i] * mul
		target[i] = target[i] - vec[i]
	}
	return target
}
