package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := scanner.Scan()
	fmt.Println(res, scanner.Text())
	var a, b int
	n, err := fmt.Scan(&a, &b)
	fmt.Print(a, b, n, err)
}
