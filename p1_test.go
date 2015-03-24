package euler

import (
    "testing"
)

func TestP1(t *testing.T) {
    
    type testpair struct {
        input int
        total int
    }

    var tests = []testpair {
        { 10, 23 },
        { 1000, 233168 },
    }

    for _, pair := range tests {
        total := P1(pair.input)
        if total != pair.total {
            t.Error(
                "\nInput:", pair.input,
                "\nExpected Total:", pair.total,
                "\nGot:", total,
            )
        }
    }
}