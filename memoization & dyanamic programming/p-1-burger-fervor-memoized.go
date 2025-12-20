package main

import "fmt"

var total_calls = 0
var cache_hit = 0
var cache_miss = 0

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func solve(m, n, t int, memo []int) int {
	total_calls++

	if memo[t] != -2 {
		cache_hit++
		return memo[t]
	}
	cache_miss++

	if t == 0 {
		memo[t] = 0
		return memo[t]
	}

	var first, second int

	if t >= m {
		first = solve(m, n, t-m, memo)
	} else {
		first = -1
	}

	if t >= n {
		second = solve(m, n, t-n, memo)
	} else {
		second = -1
	}

	if first == -1 && second == -1 {
		memo[t] = -1
		return memo[t]
	} else {
		memo[t] = max(first, second) + 1
		return memo[t]
	}

}

func solve_helper(m, n, t int) {
	memo := make([]int, t+1)
	var result int
	remaining_time := 0

	//fill the memoization array with -2
	for i := 0; i <= t; i++ {
		memo[i] = -2
	}

	for i := 0; i <= t; i++ {
		result = solve(m, n, t-i, memo)
		if result != -1 {
			remaining_time = i
			break
		}
	}

	fmt.Println(result, remaining_time)

}

func main() {
	m := 4
	n := 9
	t := 50

	solve_helper(m, n, t)

	fmt.Printf("Cache Hit: %d , Cache Miss: %d \n", cache_hit, cache_miss)
	fmt.Println("total calls: ", total_calls)
}
