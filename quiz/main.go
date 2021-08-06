package main

import (
	"fmt"
)

func couple(s string) (result []string) {
	// result := make([]string, 2)

	for s += "_"; len(s) >= 2; s = s[2:] {
		result = append(result, s[:2])
	}

	return
}

// "abcde" => []string{"ab","cd","e_"}

func main() {
	fmt.Println(couple("abcd"))
	fmt.Println(couple("abcde"))
	fmt.Println(couple("abcdef"))
}
