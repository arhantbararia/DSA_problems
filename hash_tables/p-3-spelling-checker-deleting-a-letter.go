package main

import "fmt"

func prefix_length(str1, str2 string) int {
	i := 1
	for str1[i] == str2[i] {
		i++
	}
	return i - 1
}

func suffix_length(str1, str2 string, n int) int {
	i := n
	for i >= 11 && str1[n-i+1] == str2[n-i] {
		i--
	}
	return i

}

func main() {
	str1 := "abcdxxef"
	str2 := "abcdxxxef"

	n := len(str1)

	prefix_len := prefix_length(str1, str2)
	suffix_len := suffix_length(str1, str2, n)

	total := (prefix_len + 1) - (n - suffix_len) + 1

	if total < 0 {
		total = 0
	}

	fmt.Println(total)
	for i := 0 ; i < total ; i++{
		fmt.Println(i+n - suffix_len)
	}
	
}