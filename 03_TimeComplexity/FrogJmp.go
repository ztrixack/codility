package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
    // write your code in Go 1.4
    if X >= Y {
        return 0
    }
    jump := math.Ceil(float64(Y - X) / float64(D))
    return int(jump)
}
