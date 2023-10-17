package main

import "fmt"

func GCD(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return a + b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	total := GCD(a, b)
	fmt.Println(a/total, b/total)
}
