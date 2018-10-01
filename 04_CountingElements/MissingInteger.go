package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    arr := make(map[int] bool)
    max := 0
    count := 0
    
    for _, n := range A {
        if n > 0 {
            if (arr[n] == false) {
                arr[n] = true
                count++
            }
            if max < n {
                max = n
            }
        }
    }
    
    if len(arr) == 0 {
        return 1
    }
    
    if max == count {
        return count + 1
    }
    
    for i := 1; i < count; i++ {
        if (arr[i] == false) {
            return i
        }
    }
    
    return count
}
