package euler

func P1(n int) int {
    total := 0

    for i := 0; i < n; i++ {
        if i % 3 == 0 {
            total += i
        } else if i % 5 == 0 {
            total += i
        }
    }

    return total
}