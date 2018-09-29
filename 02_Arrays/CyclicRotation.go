package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
    // write your code in Go 1.4
    if len(A) == 0 {
        return A
    }
    
    arr := append(A)
    last := len(arr) - 1
    loop := K % len(arr)
    for i := 0; i < loop; i++ {
        arr = append(arr[last:], arr[:last]...)
    }
    
    return arr
}
