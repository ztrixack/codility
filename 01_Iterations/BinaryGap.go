package solution

// you can also use imports, for example:
import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    // write your code in Go 1.4
    count, gaps := 0, 0
    bin := fmt.Sprintf("%b", N)
    for i := range bin {
        if (bin[i] == '1') {
            if (gaps < count) {
                gaps = count
            }
            count = 0
        } else if (bin[i] == '0') {
            count++
        }
    }
    return gaps
}
