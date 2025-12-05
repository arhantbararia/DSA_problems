// Input
// The first line of input is an integer n, which gives the number of snowflakes
// that we’ll be processing. The value n will be between 1 and 100,000. Each
// of the following n lines represents one snowflake: each line has six integers,
// where each integer is at least zero and at most 10,000,000.
// Output
// Our output will be a single line of text:
// • Ifthereare noidentical snowflakes, print exactly
// No two snowflakes are alike.
// • Ifthereare at least two identical snowflakes, print exactly
// Twin snowflakes found.
// The time limit for solving the test cases is two seconds.
package main

import "fmt"

const SIZE = 3

func identical_left(snow1 []int, snow2 []int, start int) bool {
	for offset := 0; offset < 6; offset++ {
		snow2index := start - offset
		if snow2index < 0 {
			snow2index += 6
		}
		if snow1[offset] != snow2[snow2index] {
			return false
		}
	}
	return true
}

func identical_right(snow1 []int, snow2 []int, start int) bool {

	for offset := 0; offset < 6; offset++ {
		if snow1[offset] != snow2[(offset+start)%6] {
			return false
		}
	}

	return true

}

func are_identical(snow1 []int, snow2 []int) bool {
	for start := 0; start < 6; start++ {
		if identical_left(snow1, snow2, start) {
			return true
		}

		if identical_right(snow1, snow2, start) {
			return true
		}
	}
	return false

}

// func identify_identical(snowflakes [][6]int, n int) {

// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if are_identical(snowflakes[i][:], snowflakes[j][:]) {
// 				fmt.Println("Twin Snowflakes found! ")
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println("No Two Snowflakes are alike")
// }

func identify_identical(snowflakes []*snowflake_node) {

	for i := 0; i < SIZE; i++ {
		node1 := snowflakes[i]
		for node1 != nil {
			node2 := node1.next
			for node2 != nil {
				if are_identical(node1.snowflake[:], node2.snowflake[:]) {
					fmt.Println("Twin Snowflakes found!")
					return
				}
				node2 = node2.next
			}
			node1 = node1.next
		}
	}
	fmt.Println("No Two Snowflakes are alike")
}

// func main1() {

// 	snowflakes1 := [][6]int{
// 		{1, 2, 3, 4, 5, 6},
// 		{7, 8, 9, 10, 11, 12},
// 		{13, 14, 15, 16, 17, 18},
// 	}
// 	identify_identical(snowflakes1, len(snowflakes1))

// 	// Test case 2: Identical snowflakes (rotated)
// 	snowflakes2 := [][6]int{
// 		{1, 2, 3, 4, 5, 6},
// 		{4, 5, 6, 1, 2, 3}, // Rotated version of the first snowflake
// 		{7, 8, 9, 10, 11, 12},
// 	}
// 	identify_identical(snowflakes2, len(snowflakes2))

// 	// Test case 3: Identical snowflakes (reversed and rotated)
// 	snowflakes3 := [][6]int{
// 		{1, 2, 3, 4, 5, 6},
// 		{6, 5, 4, 3, 2, 1}, // Reversed version of the first snowflake
// 		{1, 2, 3, 4, 5, 6}, // Exact copy of the first snowflake
// 	}
// 	identify_identical(snowflakes3, len(snowflakes3))

// 	fmt.Println("----program--over----")

// }

type snowflake_node struct {
	snowflake [6]int
	next      *snowflake_node
}

func code(snowflake [6]int) int {
	sum := 0

	for i := 0; i < 6; i++ {
		sum += snowflake[i]
	}

	return (sum % SIZE)
}

func main() {

	var snowflakes [SIZE]*snowflake_node

	// // Test case 1: No identical snowflakes
	// temp_array := [][6]int{
	// 	{1, 2, 3, 4, 5, 6},
	// 	{7, 8, 9, 10, 11, 12},
	// 	{13, 14, 15, 16, 17, 18},
	// }

	// Test case 2: Identical snowflakes (rotated)
	// temp_array := [][6]int{
	// 	{1, 2, 3, 4, 5, 6},
	// 	{4, 5, 6, 1, 2, 3}, // Rotated version of the first snowflake
	// 	{7, 8, 9, 10, 11, 12},
	// }

	// Test case 3: Identical snowflakes (reversed and rotated)
	temp_array := [][6]int{
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1}, // Reversed version of the first snowflake
		{1, 2, 3, 4, 5, 6}, // Exact copy of the first snowflake
	}

	for i := 0; i < SIZE; i++ {
		node := &snowflake_node{}

		for j := 0; j < 6; j++ {
			node.snowflake[j] = temp_array[i][j]
		}

		code := code(node.snowflake)
		fmt.Println(code)

		node.next = snowflakes[code]

		snowflakes[code] = node

	}

	identify_identical(snowflakes[:])

}
