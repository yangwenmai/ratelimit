package main

import "fmt"

func main() {
	fmt.Println("rate-limit-go")
}

// Limit 限制
func Limit() bool {
	fmt.Println("limit")
	return false
}
