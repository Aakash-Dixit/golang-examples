package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Engineers contains the list of all engineers
type Engineers struct {
	XMLName   xml.Name   `xml:"engineers"`
	Engineers []Engineer `xml:"engineer"`
}

// Engineer denotes the information available for an engineer
type Engineer struct {
	XMLName xml.Name `xml:"engineer"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// Social contains social links for engineer
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Github   string   `xml:"github"`
	Linkedin string   `xml:"linkedin"`
}

func main() {
	xmlFile, err := os.Open("engineers.xml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Opened engineers.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var engineers Engineers
	err = xml.Unmarshal(byteValue, &engineers)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(engineers.Engineers); i++ {
		fmt.Println("Engineer Type: " + engineers.Engineers[i].Type)
		fmt.Println("Engineer Name: " + engineers.Engineers[i].Name)
		fmt.Println("Github Url: " + engineers.Engineers[i].Social.Github)
	}
}
