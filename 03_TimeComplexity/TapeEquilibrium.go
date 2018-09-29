package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    arr := append(A)
    size := len(arr)
    min := 10000
    left,right := 0, 0
    for i := range arr {
        right += arr[i]
    }
    for i := 1; i < size; i++ {
        left += arr[i - 1]
        right -= arr[i - 1]
        sum := int(math.Abs(float64(left - right)))
        if (sum < min) {
            min = sum
        }
    }
    return min
}
