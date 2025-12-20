package main

import "fmt"

var total_calls = 0

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func solve(m, n, t int) int {
	total_calls++

	var first, second int

	if t == 0 {
		return 0
	}

	if t >= m {
		first = solve(m, n, t-m)
	} else {
		first = -1
	}

	if t >= n {
		second = solve(m, n, t-n)

	} else {
		second = -1
	}

	if first == -1 && second == -1 {
		return -1
	} else {
		return max(first, second) + 1
	}

}

func solve_helper(m, n, t int) {

	var result int
	remaining_time := 0

	for i := 0; i <= t; i++ {
		result = solve(m, n, t-i)
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
	t := 80

	solve_helper(m, n, t)

	fmt.Printf("Total calls to recursive function: %d", total_calls)

}
