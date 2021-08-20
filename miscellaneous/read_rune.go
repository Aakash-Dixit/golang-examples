package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	char, n, err := reader.ReadRune()
	fmt.Print(char, n, err)
	switch char {
	case 'a':
		fmt.Printf("\n%c", char)
	case 'A':
		fmt.Printf("\n%c", char)
	}
}
