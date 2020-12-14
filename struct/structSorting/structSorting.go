package main

import (
	"fmt"
	"math/rand"
	"time"
)

type student struct {
	Name string
	Roll int
}

func main() {
	// Initializing the student array with 19 elements
	arr1 := initStudent(20)
  
        // In case of large number comment the print
	fmt.Printf("\n Input array is: %v", arr1)
  
        // calling the function to sort the struct
	out := sortStruct(arr1)
  
        // In case of large number comment the print
	fmt.Printf("\n Output array is: %v", out)

}

func initStudent(limit int) []student {
	var temp student
	count := 1
	arr := []student{}
	store := make(map[int]int)
	
	source := rand.NewSource(time.Now().UnixNano())
	r :=  rand.New(source)
	
        // Infinite for loop
	for {
  
                // logic to assign unique 
                // roll number for every student
                // as random number generator can
                // generate same number
    
		roll := r.Intn(limit)
		_, ok := store[roll]
		if ok || roll == 0 {
			continue
		}else{
			store[roll] = 1
			name := fmt.Sprintf("%s%d", "student", count)
			temp.Name = name
			temp.Roll = roll
			arr = append(arr, temp)
			count = count + 1
		}
		
		if count == limit {
			break
		}
	}
	
	return arr
}

// Sort the student struct based on the Roll Number
func sortStruct(in []student) ([]student) {
	var store student
	
	// time in millisecond
	t1 := time.Now().UnixNano() / int64(time.Millisecond)
	
        // time in nanosecond
	//t1 := time.Now().UnixNano()
  
	fmt.Printf("\n t1 is: %d", t1)
	
	for i:=0; i<=len(in)-1; i++ {
		for j:=i+1; j<len(in); j++ {
			if  in[i].Roll > in[j].Roll {
				store = in[i]
				in[i] = in[j]
				in[j] = store
			}
		}	
	}
	
	// time in millisecond
	t2 := time.Now().UnixNano() / int64(time.Millisecond)
	
        // time in nanosecond
	//t2 := time.Now().UnixNano()
  
	fmt.Printf("\n t2 is: %d", t2)
  
	t3 := t2-t1
	fmt.Printf("\n Time taken to sort in millisecond is: %d", t3)
	return in
}
