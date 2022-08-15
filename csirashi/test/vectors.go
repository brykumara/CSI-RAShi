package main

import (
	"github.com/brykumara/circlclone/csidh"
)

const (
	K          = 74
	K2         = K
	Iterations = 10
	Trials     = 100
)

var rel_lat = make([]uint64, K*K)

func Sub_Multiple(target, vec []int64, mul int64) []int64 {
	for i := 0; i < csidh.PrimeCount; i++ {
		vec[i] = vec[i] * mul
		target[i] = target[i] - vec[i]
	}
	return target
}

func Reduce32(vec []float64, pool_vectors int) {
	var norm = L1NormforOneVec(vec)
	var counter = 0
	for true {
		var change = 0
		for i := 0; i < pool_vectors; i++ {
			for j := 0; j < len(pool); j++ {
				pool[j] = pool[j] + (int32)(i*K)
			}
			var plus_norm = L1NormSumforTwoVec(vec, pool)
			if plus_norm < norm {
				norm = plus_norm
				counter += 1
				AddVec(vec, pool)
				change = 1
			}
			var minus_norm = L1NormDiffforTwoVec(vec, pool)
			if minus_norm < norm {
				norm = minus_norm
				counter += 1
				SubVec(vec, pool)
				change = 1
			}
		}
		if change == 0 {
			return
		}
	}
}

func Reduce(vec []float64, pool_vectors, trials int) {
	var VEC = make([]float64, K)
	for i := 0; i < K; i++ {
		VEC[i] = vec[i]
	}
	
	var Best = make([]float64, K)
	Reduce32(VEC, pool_vectors)
	
	//memcpy(best,VEC,sizeof(int32_t)*K);
	var best_len = L1NormForOneVec(VEC)

	for j := 1; j < trials; j++ {
		//memcpy(VEC,Best,sizeof(int32_t)*K)
		for n := 0; n < 2; n++ {
			var index = // rand()%POOL_SIZE
			VEC = AddVec(VEC,(pool+K*index))
		}
		Reduce32(VEC,pool_vectors)
		norm = L1Norm(VEC)
		if norm<best_len{
			best_len = norm
			// memcpy(best,VEC,sizeof(int32_t)*K);
		}
	}

	for i := 0; i<K;i++{
		vec[i] = best[i]
	}
}

// memcpy(VEC,best,sizeof(int32_t)*K); For VEC, split Best into sizeof(int32_t)*K 
// and fit it into each element of the VEC slice. 
