package main

import (
	"fmt"
)

type person struct{
	name  string
	age   int
	address adr
}

type adr struct{
	a []string
}

func main() {
	var a person
	a.name = "test"
	a.age = 24
	a.address.a = []string{"test1", "testing"}
	//a.address.a = "test1"
	b := a
	fmt.Printf("a before changing element is: %+v\n", a)
	fmt.Printf("b before chaging element is: %+v\n", b)
	
	// if we initialize the slice, it will resolve the
	// shallow copy issue
	// b.address.a = []string{"test2", "test3"}
	b.address.a[0] = "test4"
	fmt.Printf("a after changing element is %+v\n", a)
	fmt.Printf("b after changing element is %+v\n", b)
}
