package patch

import (
	"encoding/json"
	"fmt"

	applyjsonpatch "github.com/evanphx/json-patch"
	"github.com/mattbaird/jsonpatch"
)

// Patch calculates and applies json patch
func Patch() {

	fmt.Println("######### Create json patch #########")
	jsonString1 := `{
		"users": [
		  {
			"name": "Tom",
			"type": "Engineer",
			"age": 26,
			"social": {
			  "github": "https://github.com",
			  "linkedin": "https://linkedin.com"
			}
		  },
		  {
			"name": "Sam",
			"type": "Developer",
			"age": 27,
			"social": {
			  "github": "https://github.com",
			  "linkedin": "https://linkedin.com"
			}
		  }
		]
	  }`

	jsonString2 := `{
		"users": [
		  {
			"name": "Jack",
			"type": "Tester",
			"age": 25,
			"social": {
			  "github": "https://github.com",
			  "linkedin": "https://linkedin.com"
			}
		  },
		  {
			"name": "John",
			"age": 17,
			"social": {
			  "github": "https://github.com",
			  "linkedin": "https://linkedin.com",
			  "indeed": "https://indeed.com"
			}
		  }
		]
	  }`

	jsonPatch, err := jsonpatch.CreatePatch([]byte(jsonString1), []byte(jsonString2))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, op := range jsonPatch {
		fmt.Printf("%s\n", op.Json())
	}

	jsonPatchBytes, err := json.Marshal(jsonPatch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("jsonPatch : ", string(jsonPatchBytes))

	fmt.Println("######### Apply json patch #########")

	decodedPatch, err := applyjsonpatch.DecodePatch(jsonPatchBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	modified, err := decodedPatch.Apply([]byte(jsonString1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("modified : ", string(modified))
}
