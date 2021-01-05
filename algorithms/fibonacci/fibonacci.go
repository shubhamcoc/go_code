package main

import (
	"fmt"
)
var f func(int)int
var calculation = 0

func main() {
  // Initialize the function to a variable
	result := fibonacci()
	
	fmt.Println(result(7))
	fmt.Println(result(10))
	//fmt.Println(calculation)
}

// Top function to initiate the cache variable
// We can call the function f multiple times
// without reinitializing the cache variable again.

func fibonacci() func(int) int {
	cache := make(map[int]int)
  // Closure function with dynamic programming 
	f = func(n int) int {
      // To calculate the number of times this func executed
		  calculation++
		  if _, ok := cache[n]; ok {
			    return cache[n]
		  } else {
		      if n < 2 {
			        return n
		      } else {
			        cache[n] = f(n-1) + f(n-2)
			        return cache[n]
		      }
		  }
	}
	return f
}
