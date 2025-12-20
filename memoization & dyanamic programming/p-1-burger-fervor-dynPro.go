package main

import "fmt"

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func solve(m, n, t int) {

	var result, i, first, second int
	remaining_time := 0
	dp := make([]int, t+1)

	dp[0] = 0

	for i = 1; i <= t; i++ {
		if i >= m {
			first = dp[i-m]
		} else {
			first = -1
		}

		if i >= n {
			second = dp[i-n]
		} else {
			second = -1
		}

		if first == -1 && second == -1 {
			dp[i] = -1
		} else {
			// Only add 1 if the previous state was valid (not -1)
			val1 := first
			if val1 != -1 {
				val1++
			}
			val2 := second
			if val2 != -1 {
				val2++
			}
			dp[i] = max(val1, val2)
		}

	}

	for i = 0; i <= t; i++ {
		if dp[t-i] >= 0 {
			result = dp[t-i]
			remaining_time = i
			break
		}
	}

	fmt.Println(result, remaining_time)

}

func main() {
	m := 4
	n := 9
	t := 54

	solve(m, n, t)

}
