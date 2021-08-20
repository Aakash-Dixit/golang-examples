package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start reading from shell...")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\r\n", "")
		fmt.Println(text)
	}

}
