package main

import (
	"fmt"
	"encoding/json"
)

// Complex Struct which contain value and reference type data structure.
// For Ex:
// int, float, string, struct are pass by value
// slice, map these are pass by reference
type person struct{
	// Composite Literals (i.e.) json is required to convert the struct in Json 
	// It work with var first letter as capital
	// (i.e.) with exported filed 
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Alias []string `json:"alias"`
	Address adr  `json:"address"`
	SomeDetail  map[string]interface{}  `json:"somedetail"`
}

type adr struct{
	Line1 []string  `json:"line1"`
	Number num      `json:"number"`
}

type num struct{
	Number1 string	`json:"number1"`
}

// Struct without reference type data structure
type person1 struct{
	// without composite literals and Exported fields
	name  string 
	age   int    
	address adr1  
}

type adr1 struct{
	line1 string  
	city string
}

func main() {
	// initialize the struct person
	var a person
	a.Name = "test"
	a.Age = 24
	a.Alias = []string{"alias1", "alias2"}
	a.SomeDetail = map[string]interface{}{"key1": "value1"}
	a.Address.Line1 = []string{"line1", "line2"}
	a.Address.Number.Number1 = "123456789"

	// shallow copy of the struct
	b := a
	fmt.Println("In case of Shallow Copy of complex struct like person")
	fmt.Println()
	fmt.Println("Before changing element in object b")
	fmt.Printf("a is: %+v\n", a)
	fmt.Printf("b is: %+v\n", b)
	fmt.Println()
	
	// if we initialize the slice, it will resolve the
	// shallow copy issue
	// b.address.Line1 = []string{"test2", "test3"}
	
	// Changing the element in the copy object b
	// In case of deep copy it should not affect the 
	// object a, but this will change the object a.
	b.Address.Line1[0] = "test4"
	fmt.Println("After changing element (i.e.) line1 = test4 in object b")
	fmt.Printf("a is %+v\n", a)
	fmt.Printf("b is %+v\n", b)
	fmt.Println()
	
	// Calling the deepCopy func to create a new object c
	c := deepCopy(a)
	
	// Changing the element in the copy object c
	// In case of deep copy it should not affect the 
	// object a, and it will not change the object a.
	c.Address.Line1[0] = "test5"
	c.SomeDetail = map[string]interface{}{"key2": "value2"}
	fmt.Println("In case of Deep Copy")
	fmt.Println()
	fmt.Println("After changing element (i.e.) test4 = test5, key1 = key2 and value1 = value2 in object c")
	fmt.Printf("a is %+v\n", a)
	fmt.Printf("c is %+v\n", c)
	fmt.Println()
	
	fmt.Println("In case of Shallow Copy of struct like person1")
	fmt.Println()
	// initialize the struct person
	var per1 person1
	per1.name = "person1"
	per1.age = 26
	per1.address.line1 = "line64"
	per1.address.city = "somecity"
	
	// shallow copy of the struct
	per2 := per1
	fmt.Println("Before changing element in object per2")
	fmt.Printf("per1 is: %+v\n", per1)
	fmt.Printf("per2 is: %+v\n", per2)
	fmt.Println()
	
	// Changing the element in the copy object per2
	// In this case it will not change the object per1
	// as we don't have reference data structure in person1.
	per2.address.line1 = "line54"
	fmt.Println("After changing element (i.e.) line64 = test54 in object per2")
	fmt.Printf("per1 is %+v\n", per1)
	fmt.Printf("per2 is %+v\n", per2)
	fmt.Println()
	
	// shallow copy of the struct
	per3 := per1
	fmt.Println("Before changing element in object per3")
	fmt.Printf("per1 is: %+v\n", per1)
	fmt.Printf("per2 is: %+v\n", per2)
	fmt.Printf("per3 is: %+v\n", per3)
	fmt.Println()
	
	// Changing the element in the copy object per2
	// In this case it will not change the object per1
	// as we don't have reference data structure in person1.
	per3.address.line1 = "line44"
	fmt.Println("After changing element (i.e.) line64 = test44 in object per3")
	fmt.Printf("per1 is %+v\n", per1)
	fmt.Printf("per2 is %+v\n", per2)
	fmt.Printf("per3 is %+v\n", per3)
	fmt.Println()
}

func copy(nested interface{}, copyMap map[string]interface{}) {
        
	switch nested.(type){
	case map[string]interface{}:
		for k, v := range nested.(map[string]interface{}){
			copyMap[k] = v
		}
	}			
}

func deepCopy(p person) person {
	data := make(map[string]interface{})
	tempData := make(map[string]interface{})
	var b person
	
	js, _ := json.Marshal(p)
	
	err := json.Unmarshal(js, &data)
	if err != nil{
		fmt.Println("err is:", err)
	}
	
	copy(data, tempData)
	//fmt.Printf("tempData is: %+v\n", tempData)
	
	js1, _ := json.Marshal(tempData)
	err = json.Unmarshal(js1, &b)
	if err != nil{
		fmt.Printf("Error is: %+v\n", err)
	}
	return b
}
