package main

import (
	"fmt"
	"os/exec"
)

func execute() {
	op1, err := exec.Command("dir").Output()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(op1[:]))

	op2, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(op2[:]))

	/*
		op3, err := exec.Command("ls","-ltr").Output()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(op3[:]))

		build with : env GOOS=linux GOARCH=amd64 go build syatem_exec.go
	*/
}

func main() {
	/*if runtime.GOOS == "windows" {
		fmt.Print("Cannot perform op for windows")
		return
	}*/
	execute()
}
