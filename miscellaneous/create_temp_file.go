package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tempDir, err := ioutil.TempDir(`C:\Users\user\Desktop\golang-examples\miscellaneous`, "mydir-")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.RemoveAll(tempDir)
		if err != nil {
			log.Println(err)
		}
	}()
	file, err := ioutil.TempFile(tempDir, "myfile-*.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.Remove(file.Name())
		if err != nil {
			log.Println(err)
		}
	}()

	defer file.Close()
	log.Println("Temp file name : ", file.Name())

	if _, err := file.Write([]byte("hello world\n")); err != nil {
		log.Println(err)
		return
	}

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("File Data : ", string(data))
}
