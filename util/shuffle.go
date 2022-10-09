package util

import (
	"math/rand"
)

// Shuffle 洗牌算法,用于打乱染色体
func Shuffle(x []float64) []float64 {
	res := make([]float64, len(x))
	copy(res, x)
	n := len(res)
	for i := range res {
		j := i + rand.Intn(n-i)
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func ShuffleFloat(x []float64) []float64 {
	res := make([]float64, len(x))
	copy(res, x)
	//rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}
func ShuffleInt(x []int) []int {
	res := make([]int, len(x))
	copy(res, x)
	//rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

// RandFloat64 1e(-9)级别的返回随机小数
func RandFloat64() float64 {
	return float64(rand.Intn(1e16)) / 1e16
}
