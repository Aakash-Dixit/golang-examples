package marshal

import (
	"encoding/json"
	"fmt"
	"os"
)

// Student represents a student's info
type Student struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subject"`
}

// Subject represents a subject's details
type Subject struct {
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

// Marshal demonstrates json marshal ops
func Marshal() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(5.55)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("golang-examples")
	fmt.Println(string(strB))

	slcD := []string{"golang", "docker", "kubernetes"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"helm": 5, "etcd": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	subjects := []Subject{
		{Name: "physics", Marks: 80},
		{Name: "maths", Marks: 85},
	}
	student := Student{Name: "John", Subjects: subjects}

	byteArray, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteArray))

	byteArray, err = json.MarshalIndent(student, "", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteArray))

	mapJSON := map[string]interface{}{
		"name": "aak",
		"age":  26,
	}
	byteArray, err = json.Marshal(mapJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteArray))

	fmt.Println("###################### encode json ####################")

	json.NewEncoder(os.Stdout).Encode(student) //can be done for map[string]interface{}
}
