package dynamic

import "golang.org/x/tools/container/intsets"

var cache map[int]int = map[int]int{}

func solve(n int, coins []int) int {
	cache[0] = 0
	for i := 1; i <= n; i++ {
		cache[i] = intsets.MaxInt
	}

	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if coin <= i {
				cache[i] = min(cache[i], 1+cache[i-coin])
			}
		}
	}

	return cache[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
