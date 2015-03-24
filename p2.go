package euler

/*
    Returns fibonacci value
    Param n is nth fibonacci term
    Param memtable is the memoization table, storing all fib terms
*/
func Fib(n int, memtable map[int]int) int {

    if n <= 1 {
        return 1
    }

    if term, ok := memtable[n]; ok {
        return term
    }

    fib_1 := Fib(n - 1, memtable)
    fib_2 := Fib(n - 2, memtable)

    memtable[n] = fib_1 + fib_2

    return fib_1 + fib_2
}

/*
    Returns nth fib term
    Param n is nth fibonacci term
    Param limit is the maximum number to take in as a sum for even terms
*/
func P2(n int, limit int) int {

    memtable := map[int]int{}
    Fib(n, memtable)

    evensum := 0

    for _, term := range memtable {
        if term <= limit && term % 2 == 0 {
            evensum += term
        }
    }

    return evensum
}