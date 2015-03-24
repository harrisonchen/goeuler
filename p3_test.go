package euler

import (
    "testing"
)

func TestIsPrime(t *testing.T) {
    type testpair struct {
        num int
        primes []int
        result bool
    }

    var tests = []testpair {
        { 0, []int{}, false },
        { 1, []int{}, true },
        { 2, []int{}, true },
        { 3, []int{}, true },
        { 4, []int{}, false },
        { 5, []int{}, true },
        { 6, []int{}, false },
        { 7, []int{}, true },
        { 9, []int{}, false },
        { 19, []int{}, true },
        { 121, []int{}, false },
        { 181, []int{}, true },
        { 191, []int{}, true },
        { 289, []int{}, false },
        { 36481, []int{}, false },
    }

    for _, pair := range tests {
        result := isPrime(pair.num, &pair.primes)
        if result != pair.result {
            t.Error(
                "-> Input:", pair.num,
                "-> Expected:", pair.result,
                "-> Got:", result,
            )
        }
    }
}

// This is also the same tests for P3
func TestMaxFactor(t *testing.T) {
    type testpair struct {
        num int
        maxPrime int
    }

    var tests = []testpair {
        { 1, 1 },
        { 2, 2 },
        { 3, 3 },
        { 4, 2 },
        { 5, 5 },
        { 6, 3 },
        { 9, 3 },
        { 19, 19 },
        { 121, 11 },
        { 191, 191 },
        { 289, 17 },
        { 36481, 191 },
        { 600851475143, 6857 },
    }

    for _, pair := range tests {
        maxPrime := maxPrimeFactor(pair.num)
        if maxPrime != pair.maxPrime {
            t.Error(
                "-> Input:", pair.num,
                "-> Expected:", pair.maxPrime,
                "-> Got:", maxPrime,
            )
        }
    }
}