package euler

import (
    "testing"
)

func TestLargestPalindrome(t *testing.T) {
    type testpair struct {
        digits int
        palindrome int
    }

    var tests = []testpair {
        { 2, 9009 },
    }

    for _, pair := range tests {
        palindrome := largestPalindrome(pair.digits)
        if palindrome != pair.palindrome {
            t.Error(
                "-> Input:", pair.digits,
                "-> Expected:", pair.palindrome,
                "-> Got:", palindrome,
            )
        }
    }
}