package main

import "fmt"

var num int = 5

func init() {
	fmt.Println("Inside init")
	fmt.Println(num)
	num = 10
	fmt.Println(num)
}

func main() {
	fmt.Println("Inside main")
	fmt.Println(num)
}
