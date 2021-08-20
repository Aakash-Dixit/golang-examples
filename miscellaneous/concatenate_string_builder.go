package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString strings.Builder
	myString.WriteString("Hello")
	fmt.Println(myString.String())
	myString.WriteString("World")
	fmt.Println(myString.String())
}
