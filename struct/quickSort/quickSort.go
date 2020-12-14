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
	
	arr1 := initStudent(20)
	fmt.Printf("\n Input array is: %v", arr1)
	
	t1 := time.Now().UnixNano() / int64(time.Millisecond)
	quickSort(arr1)
	t2 := time.Now().UnixNano() / int64(time.Millisecond)
	t3 := t2 - t1
	fmt.Printf("\n Time take to sort the Struct in millisecond: %d", t3)
        fmt.Printf("\n Output array is: %v", arr1)

}

func initStudent(limit int) []student {
	var temp student
	count := 1
	arr := []student{}
	store := make(map[int]int)
	
	source := rand.NewSource(time.Now().UnixNano())
	r :=  rand.New(source)
	
	for {
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

func quickSort(arr []student) {
    size := len(arr)
    
    if size < 2 {
        return 
    }
    
    left := 0
    right := size - 1
    
    for i, _ := range arr {
        if arr[i].Roll < arr[right].Roll {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
    arr[left], arr[right] = arr[right], arr[left]
    
    quickSort(arr[:left])
    quickSort(arr[left+1:])    
}
