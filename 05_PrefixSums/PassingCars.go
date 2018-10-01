package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    count := 0
    multiply := 0
    
    for _, car := range A {
        if car == 0 {
            multiply++
        } 
        if multiply > 0 {
            if car == 1 {
                count += multiply
                if count > 1000000000 {
                    return -1
                }
            }
        }
    }
    
    return count
}
