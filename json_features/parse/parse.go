package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// Engineers contains the list of all engineers
type Engineers struct {
	Engineers []Engineer `json:"engineers"`
}

// Engineer denotes the information available for an engineer
type Engineer struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Social contains social links for engineer
type Social struct {
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
}

// Parse demonstrates json parse operations
func Parse() {
	jsonFile, _ := os.Open("engineers.json")
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	//jsonBytes, _ := ioutil.ReadFile("engineers.json")
	var engineers Engineers
	json.Unmarshal(jsonBytes, &engineers)
	fmt.Print(engineers)

	for i := 0; i < len(engineers.Engineers); i++ {
		fmt.Println("Engineer Type: " + engineers.Engineers[i].Type)
		fmt.Println("Engineer Age: " + strconv.Itoa(engineers.Engineers[i].Age))
		fmt.Println("Engineer Name: " + engineers.Engineers[i].Name)
		fmt.Println("Github Url: " + engineers.Engineers[i].Social.Github)
	}

	fmt.Println("###################### Reading with map[string]interface{} ####################")
	var jsonObj map[string]interface{}
	err := json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	engineersObj := jsonObj["engineers"].([]interface{})
	fmt.Printf("Type: %T Value: %v \n", engineersObj, engineersObj)
	for _, engineerObj := range engineersObj {
		engineerObjMap := engineerObj.(map[string]interface{})
		name := engineerObjMap["name"].(string)
		fmt.Println(name)
		age := engineerObjMap["age"].(float64)
		fmt.Println(age)
		socialMap := engineerObjMap["social"].(map[string]interface{})
		github := socialMap["github"].(string)
		linkedin := socialMap["linkedin"].(string)
		fmt.Println(github)
		fmt.Println(linkedin)
	}
	fmt.Println("###################### Reading with gjson ####################")
	jsonString := string(jsonBytes)
	if !gjson.Valid(jsonString) {
		fmt.Println("Invalid json")
	} else {
		fmt.Println("Valid json")
	}
	fmt.Println(jsonString)
	lenEngineers := gjson.Get(jsonString, "engineers.#").Int()
	for i := 0; i < int(lenEngineers); i++ {
		index := "engineers." + strconv.Itoa(i) + "."
		name := gjson.Get(jsonString, index+"name")
		fmt.Println(name)
		typeEngineer := gjson.Get(jsonString, index+"type")
		fmt.Println(typeEngineer)
		age := gjson.Get(jsonString, index+"age")
		fmt.Println(age)
		github := gjson.Get(jsonString, index+"social.github")
		fmt.Println(github)
		linkedin := gjson.Get(jsonString, index+"social.linkedin")
		fmt.Println(linkedin)

		indeed := gjson.Get(jsonString, index+"social.indeed")
		if indeed.Exists() {
			fmt.Println(indeed)
		} else {
			fmt.Println("indeed does not exist")
		}
	}

	//iterating
	result := gjson.Get(jsonString, "engineers.#.name")
	for _, name := range result.Array() {
		fmt.Println(name.String())
	}

	//querying
	name := gjson.Get(jsonString, `engineers.#(name="Roman").type`)
	fmt.Println(name.String()) // prints "Elliot"

	//getting many
	results := gjson.GetMany(jsonString, "engineers.0.name", "engineers.1.age")
	for _, val := range results {
		fmt.Println(val)
	}

	//working with jsonBytes
	result = gjson.GetBytes(jsonBytes, "engineers.0.social.github")
	fmt.Println(string([]byte(result.Raw)))

	fmt.Println("###################### decode json ####################")
	var engineers1 Engineers
	reader := strings.NewReader(string(jsonBytes))
	json.NewDecoder(reader).Decode(&engineers1)
	fmt.Println(engineers1)
}
