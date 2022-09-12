package main

import (
	c "crypto/rand"
	"math/big"

	"github.com/brykumara/circlclone/csidh"
)

const (
	K          = 74
	K2         = K
	Iterations = 10
	Trials     = 100
)

var rel_lat = make([]float64, K*K)

func Sub_Multiple(target, vec []float64, mul float64) []float64 {
	for i := 0; i < csidh.PrimeCount; i++ {
		target[i] -= vec[i] * mul
	}
	return target
}

func Reduce32(vec []float64, pool_vectors float64) {
	var norm = L2NormforOneVec(vec) // CHANGE
	var counter = 0
	for {
		var change = 0
		for i := 0; i < (int)(pool_vectors); i++ {
			for j := 0; j < len(pool); j++ {
				pool[j] = pool[j] + (int32)(i*K)
			}
			poolconv := make([]float64, len(pool))
			for n := 0; n < len(poolconv); n++ {
				poolconv[n] = (float64)(pool[n])
			}
			var plus_norm = L2NormSumforTwoVec(vec, poolconv) // CHANGE
			if plus_norm < norm {
				norm = plus_norm
				counter += 1
				AddVec(vec, poolconv)
				change = 1
			}
			var minus_norm = L2NormDiffforTwoVec(vec, poolconv) //CHANGE
			if minus_norm < norm {
				norm = minus_norm
				counter += 1
				SubVec(vec, poolconv)
				change = 1
			}
		}
		if change == 0 {
			return
		}
	}
}

func Reduce(vec []float64, pool_vectors, trials float64) {
	var VEC = make([]float64, K)
	for i := 0; i < K; i++ {
		VEC[i] = vec[i]
	}

	var Best = make([]float64, K)
	Reduce32(VEC, pool_vectors)

	Best = VEC
	var best_len = L2NormforOneVec(VEC) // CHANGE

	for j := 1; j < (int)(trials); j++ {
		VEC = Best
		for n := 0; n < 2; n++ {
			index, _ := c.Int(rng, big.NewInt((int64)(rand_max)))
			var num = index.Int64() % 10000
			newpool := make([]float64, len(pool))
			for k := 0; k < len(pool); k++ {
				newpool[k] = (float64)(pool[k])
				newpool[k] = newpool[k] + (float64)(74*num)
			}
			VEC = AddVec(VEC, newpool)
		}
		Reduce32(VEC, pool_vectors)
		var norm = L2NormforOneVec(VEC) // CHANGE
		if norm < best_len {
			best_len = norm
			Best = VEC
		}
	}

	for i := 0; i < K; i++ {
		vec[i] = Best[i]
	}
}
