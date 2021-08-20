package main

import (
	"errors"
	"fmt"
)

//Set represents the set datastucture
type Set struct {
	Elements map[string]struct{}
}

//NewSet returns a new set
func NewSet() *Set {
	set := Set{
		Elements: make(map[string]struct{}),
	}
	return &set
}

//Add add elements to a struct
func (set *Set) Add(elem string) {
	set.Elements[elem] = struct{}{}
}

//Delete deletes an element from the set if it exists
func (set *Set) Delete(elem string) error {
	if set.Contains(elem) {
		delete(set.Elements, elem)
		return nil
	}
	return errors.New("element : " + elem + ", not found in set")
}

//Contains checks if an element exists in the set
func (set *Set) Contains(elem string) bool {
	_, exists := set.Elements[elem]
	return exists
}

//List lists the key(name) of the ele
func (set *Set) List() {
	for key := range set.Elements {
		fmt.Println(key)
	}
}

func main() {
	fmt.Println("Start Set Operations")
	set := NewSet()

	set.Add("apple")
	set.Add("orange")
	set.Add("grapes")
	set.List()

	err := set.Delete("grapes")
	if err != nil {
		fmt.Println(err)
	}
	set.List()

	err = set.Delete("banana")
	if err != nil {
		fmt.Println(err)
	}
}
