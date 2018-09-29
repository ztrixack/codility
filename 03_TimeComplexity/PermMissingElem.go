package solution

// you can also use imports, for example:
// import "fmt"
import "sort"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    arr := append(A)
    size := len(arr) - 1
    sort.Ints(arr)
    
    for i := 0; i < size; i++ {
        if arr[i] + 1 != arr[i + 1] {
            return  arr[i] + 1
        }
    }
    
    if (size < 0) {
        return 1
    } else if (arr[0] != 1) {
        return arr[0] - 1
    } else {
        return arr[size] + 1
    }
}
