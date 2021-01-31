package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    arr := initArray(10)
    arr1 := initArray(10)
    t1 := time.Now().Nanosecond()
    result := selectionSort(arr)
    t2 := time.Now().Nanosecond()
    
    t3 := time.Now().Nanosecond()
    result1 := modifiedSelectionSort(arr1)
    t4 := time.Now().Nanosecond()
    
    fmt.Println(result)
    fmt.Println("time taken by selection sort:", t2-t1)
    fmt.Println(result1)
    fmt.Println("time taken by modified selection sort:", t4-t3)
}

func initArray(limit int) []int {
	count := 1
	var arr []int
	store := make(map[int]int)
	
	source := rand.NewSource(time.Now().UnixNano())
	r :=  rand.New(source)
	
    // Infinite for loop
	for {
  
        // logic to assign unique 
        // roll number for every student
        // as random number generator can
        // generate same number
    
		roll := r.Intn(limit+100)
		_, ok := store[roll]
		if ok || roll == 0 {
			continue
		}else{
			store[roll] = 1
			arr = append(arr, roll)
			count = count + 1
		}
		
		if count == limit {
			break
		}
	}
	
	return arr
}

func selectionSort(A []int) []int {
	size := len(A)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i+1; j < size; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		
		if min != i {
			A[i], A[min] = A[min], A[i]
		}
	}
	
	return A
}

func modifiedSelectionSort(A []int) []int {
    size := len(A)
	for i := 0; i < size-1; i++ {
	    arr := A[i:]
	    size := len(arr)
		  min := arr[0]
		  max := arr[size-1] 
		  p := 0
	    low := 0
	    if low != size -1 {
	    _, min, p = findMaxMin(arr, low, size-1, max, min, p)
	    }
	    if i != p+i {
	        A[i], A[p+i] = A[p+i], A[i]
	    }
    }  
    
    return A
}

func findMaxMin(arr []int, low, high, max, min, p int) (int, int, int) {
	// If only one element is present then
	// compare the max with it to check if max
	// is more or the element present in the array
	// similarly check if min is lower than the
	// element and return the max and min
	if low == high {
		if max < arr[high] {
			max = arr[high]
		}
		if min > arr[low] {
			min = arr[low]
			p = low
		}
		return max, min, p
	}
	
	if high - low == 1 {
		if arr[high] > arr[low] {
			if max < arr[high] {
				max = arr[high]
			}
			
			if min > arr[low] {
				min = arr[low]
				p = low
			}
		} else {
			if max < arr[low] {
				max = arr[low]
			}
			
			if min > arr[high] {
				min = arr[high]
				p = high
			}
		}
		
		return max, min, p
	}
	
	mid := (high + low) / 2
	max, min, p = findMaxMin(arr, low, mid, max, min, p)
	max, min, p  = findMaxMin(arr, mid+1, high, max, min, p)
	
	return max, min, p
}
