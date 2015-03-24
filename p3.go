package euler

import (
    "strconv"
)

/*
    Checks if number is prime
    Param n is the number to be checked
    Param primes is a pointer to slice of primes up to the input number
    *** Totally unneeded for Euler Problem #3 ***
*/
func isPrime(n int, primes *[]int) bool {

    if n == 0 {
        return false
    }

    if n == 1 {
        *primes = append(*primes, 1)
        return true
    }

    if n == 2 {
        *primes = append(*primes, 1, 2)
        return true
    }

    if n == 3 {
        *primes = append(*primes, 1, 2, 3)
        return true
    }

    *primes = append(*primes, 2, 3)

    for i := 4; i <= n; i++ {
        prime := true
        for j := 0; j < len(*primes); j++ {
            if i % (*primes)[j] == 0 {
                prime = false
                break;
            }
        }

        if prime {
            *primes = append(*primes, i)
        }
    }

    if (*primes)[len(*primes) - 1] == n {
        return true
    }

    return false;
}

/*
    Recursive helper function to find max prime factor
    Param n is the number to be check for factors
    *** NEEDS REFACTORING ***
*/
func maxPrimeFactorHelper(n int) int {

    if n == 1 {
        return 1
    }

    if n == 2 {
        return 2
    }

    if n == 3 {
        return 3
    }

    if n == 5 {
        return 5
    }

    maxFactor := 1

    for ;; {

        divides := false
        if n % 2 == 0 {
            n = n / 2
            if 2 > maxFactor {
                maxFactor = 2
            }
            divides = true
        }
        if n % 3 == 0 {
            n = n / 3
            if 3 > maxFactor {
                maxFactor = 3
            }
            divides = true
        }
        if n % 5 == 0 {
            n = n / 5
            if 5 > maxFactor {
                maxFactor = 5
            }
            divides = true
        }

        if n <= 1 {
            return maxFactor
        }

        if !divides {
            break
        }
    }

    length := len(strconv.Itoa(n))

    length = length / 2 + 1

    newDomainString := ""

    for i := 0; i < length; i++ {
        newDomainString += "9"
    }

    newDomain, _ := strconv.Atoi(newDomainString)

    maxFactor = -1

    for i := 7; i < newDomain; i++ {
        if n % i == 0 {
            if n == i {
                return n
            }

            factor_1 := maxPrimeFactorHelper(n % i)
            factor_2 := maxPrimeFactorHelper(i)

            if factor_1 > factor_2 {
                maxFactor = factor_1
            } else {
                maxFactor = factor_2
            }
        }
    }

    if maxFactor == -1 {
        maxFactor = n
    }

    return maxFactor
}

/*
    Finds the max prime factor of a number
    Param n is the input number to be factored
*/
func maxPrimeFactor(n int) int {
    return maxPrimeFactorHelper(n)
}

func P3(n int) int {
    return maxPrimeFactor(n)
}