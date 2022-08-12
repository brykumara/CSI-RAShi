package main

import (
	"github.com/brykumara/circlclone/csidh"
)

func Sub_Multiple(target, vec []int64, mul int64) []int64 {
	for i := 0; i < csidh.PrimeCount; i++ {
		vec[i] = vec[i] * mul
		target[i] = target[i] - vec[i]
	}
	return target
}
