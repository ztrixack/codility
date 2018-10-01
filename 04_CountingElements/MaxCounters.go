package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
    // write your code in Go 1.4
    max, lastMax := 0, 0
    counters := make([]int, N)
    
    for _, n := range A {
        if n >= 1 && n <= N {
            if lastMax > counters[n - 1] {
                counters[n - 1] = lastMax
            }
            counters[n - 1]++
            if max < counters[n - 1] {
                max = counters[n - 1]
            }
        } else if (n == N + 1) {
            lastMax = max
        }
    }
    
    for i := range counters {
        if counters[i] < lastMax {
            counters[i] = lastMax
        }
    }
    
    return counters
}
