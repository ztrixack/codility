package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type dna struct {
    A int
    C int
    G int
    T int
}

func Solution(S string, P []int, Q []int) []int {
    // write your code in Go 1.4
    sizeString := len(S) + 1
    sizeResult := len(P)
    dnas := make([]dna, sizeString)
    for i := 1; i < sizeString; i++ {
        dnas[i].A = dnas[i - 1].A
        dnas[i].C = dnas[i - 1].C
        dnas[i].G = dnas[i - 1].G
        dnas[i].T = dnas[i - 1].T
        
        if S[i - 1] == "A"[0] {
            dnas[i].A++
        } else if S[i - 1] == "C"[0] {
            dnas[i].C++
        } else if S[i - 1] == "G"[0] {
            dnas[i].G++
        } else {
            dnas[i].T++
        }
    }
    
    var result = make([]int, sizeResult)
    for i := 0; i < sizeResult; i++ {
        from, to := P[i], Q[i] + 1
        if dnas[to].A - dnas[from].A > 0  {
            result[i] = 1
        } else if dnas[to].C - dnas[from].C > 0 {
            result[i] = 2
        } else if dnas[to].G - dnas[from].G > 0 {
            result[i] = 3
        } else {
            result[i] = 4
        } 
    }
    
    return result
}
