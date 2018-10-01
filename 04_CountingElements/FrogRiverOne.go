package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
    // write your code in Go 1.4
    bucket := make([]bool, X + 1)
    count := 0
    
    for i, n := range A {
        if (bucket[n] == false) {
            bucket[n] = true
            count++
        }
        
        if (count == X) {
            return i
        }
	}
	
	return -1
}
