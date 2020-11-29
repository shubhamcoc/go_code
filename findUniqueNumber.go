package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {
	// Generate a new seed value
	// so that random num generator return 
	// new random number everytime or else
	// random number generator will
	// generate same number again and again
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var arr []int
	for i := 0; i<1000000; i++{
		num := r.Intn(10)
		arr = append(arr, num)
	}
    // do not uncomment as this will print the array
    // with 1 million elements
	// fmt.Println(arr)
  
	result := uniqueWithMap(arr)
    // unsorted result, we will get the unique 
    // numbers but not in increasing order
    // taking less time than using sort algo
    // Eg: [2, 4, 1, 0, 9]
	fmt.Println(result)	
	
	results := uniqueWithSort(arr)
    // sorted result, this is taking more time than
    // map algo
    // Eg: [0, 1, 2, 4, 7]
	fmt.Println(results)
	
}

// finding unique numbers in large array
// using map data structure
func uniqueWithMap(arr []int) []int {
unique := make(map[int]int)
	count := 0
	t1 := time.Now().UnixNano()
	for _, val := range arr {
		count = count+1
		_, ok := unique[val]
		if ok{
            // If we uncomment below line
            // this algo is taking more time 
            // than sorting
			// unique[val] = unique[val] + 1
			continue
		}else{
			unique[val] = 1
		}		
	}
	var finalarr []int
	for k, _ := range unique{
		finalarr = append(finalarr, k)
	}
	t2 := time.Now().UnixNano()
	fmt.Println("t1 in uniqueWithMap is:", t1)
	fmt.Println("t2 in uniqueWithMap is:", t2)
	t3 := t2 - t1
	t3 = t3 / int64(time.Millisecond)
	fmt.Printf("time taken in uniqueWithMap is: %+v\n", t3)
	return finalarr
}

// finding unique numbers in large array
// using sort function
func uniqueWithSort(arr []int) []int {
	var finalarr []int
	t1 := time.Now().UnixNano() 
	sort.Ints(arr)
	for i := 0; i<len(arr)-1; i++{
		if arr[i] != arr[i+1]{
			finalarr = append(finalarr, arr[i])
		}
		if i == (len(arr)-2){
			finalarr = append(finalarr, arr[i])
		}
	}
	t2 := time.Now().UnixNano()
	fmt.Println("t1 in uniqueWithSort is:", t1)
	fmt.Println("t2 in uniqueWithSort is:", t2)
	t3 := t2 - t1
	t3 = t3 / int64(time.Millisecond)
	fmt.Printf("time taken in uniqueWithSort is: %+v\n", t3)
	return finalarr
}
