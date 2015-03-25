package euler

import (
    "fmt"
    "strconv"
)

func maxNumberWithDigits(digits int) int {
    maxNumberString := ""

    for i := 0; i < digits; i++ {
        maxNumberString += "9"
    }

    maxNumber, _ := strconv.Atoi(maxNumberString)
    return maxNumber
}

func isPalindrome(n int) bool {
    numberString := strconv.Itoa(n)

    length := len(numberString)
    upperHalf := numberString[:(length / 2)]
    lowerHalf := numberString[(length / 2 + length % 2):]

    if upperHalf == reverseString(lowerHalf) {
        return true
    }
    return false;
}

// Returns next lowest palindrome
func reducePalindrome(n int) int {
    numberString := strconv.Itoa(n)
    length := len(numberString)
    midlen := length / 2
    reduceValueString := ""

    if length % 2 == 1 {
        reduceValueString += "1"
        for i := midlen + 1; i < length; i++ {
            reduceValueString += "0"
        }
    } else {
        reduceValueString += "11"
        for i := midlen + 1; i < length; i++ {
            reduceValueString += "0"
        }
    }
    reduceValue, _ := strconv.Atoi(reduceValueString)

    // if length % 2 == 1 {
    //     carryString := ""
    //     if string(numberString[midlen]) == "0" {
    //         string(numberString[midlen]) = "9"
    //     }
    // } else {

    // }

    return reduceValue
}

// Chops the ends of the palindrome
func chopPalindrome(n int) int {
    numberString := strconv.Itoa(n)
    newString := ""
    for i := 1; i < len(numberString) - 1; i++ {
        newString += string(numberString[i])
    }
    newPalindrome, _ := strconv.Atoi(newString)
    return newPalindrome
}

func reverseString(s string) string {
    reversed := ""
    for i := len(s) - 1; i >= 0; i-- {
        reversed += string(s[i])
    }

    return reversed
}

func largestPalindrome(n int) int {
    fmt.Println(maxNumberWithDigits(n))
    fmt.Println(isPalindrome(9009))
    fmt.Println(reverseString("123"))
    fmt.Println(chopPalindrome(123321))
    fmt.Println(reducePalindrome(999999))
    fmt.Println(reducePalindrome(99999))
    fmt.Println(reducePalindrome(9999))
    return -1
}