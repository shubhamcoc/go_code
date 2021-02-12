package main

import (
	"fmt"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	c3 := 0
	c5 := 0
	for i := 1; i < 101; i++ {
		c3++
		c5++
		w := ""
		if c3 == 3 { w += "fizz"; c3 = 0 }
		if c5 == 5 { w += "buzz"; c5 = 0 }
		
		if w != "" {
			fmt.Println(w)
		}else{
			fmt.Println(i)
		}
	}
}
