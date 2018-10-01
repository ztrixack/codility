package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    arr := append(A)
    size := len(arr) - 1
    sort.Ints(arr)
    
    if (arr[0] != 1) {
        return 0
    }
    
    for i := 0; i < size; i++ {
        if arr[i] + 1 != arr[i + 1] {
            return  0
        }
    }
    
    return 1
}
