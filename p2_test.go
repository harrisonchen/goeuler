package euler

import (
    "testing"
)

func TestFib(t *testing.T) {

    type testpair struct {
        input int
        fib int
        memtable map[int]int
    }

    var tests = []testpair {
        { 0, 1, map[int]int{} },
        { 1, 1, map[int]int{} },
        { 2, 2, map[int]int{} },
        { 3, 3, map[int]int{} },
        { 10, 89, map[int]int{} },
        { 33, 5702887, map[int]int{} },
    }

    for _, pair := range tests {
        fib := Fib(pair.input, pair.memtable)
        if fib != pair.fib {
            t.Error(
                "Input:", pair.input,
                "Expected:", pair.fib,
                "Got:", fib,
            )
        }
    }
}

func TestP2(t *testing.T) {

    type testpair struct {
        input int
        limit int
        sum int
    }

    var tests = []testpair {
        { 0, 4000000, 0 },
        { 1, 4000000, 0 },
        { 2, 4000000, 2 },
        { 3, 4000000, 2 },
        { 10, 4000000, 44 },
        { 33, 4000000, 4613732 },
    }

    for _, pair := range tests {
        sum := P2(pair.input, pair.limit)
        if sum != pair.sum {
            t.Error(
                "Input:", pair.input,
                "Expected:", pair.sum,
                "Got:", sum,
            )
        }
    }
}