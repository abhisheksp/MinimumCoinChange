package minimumcoinchange

import "golang.org/x/tools/container/intsets"

func solve(n int, coins []int) int {
	if n == 0 {
		return 0
	}

	minVal := intsets.MaxInt
	for _, coin := range coins {
		if coin <= n {
			minVal = min(minVal, solve(n-coin, coins)+1)
		}
	}

	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
