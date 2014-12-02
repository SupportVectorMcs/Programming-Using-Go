// ANDREW ID: longh
// NAME: Long He

package main

import (
    "fmt"
    "math"
)

/*==============================================================================
 * 1. Gauss' formula.
 *
 * Implement a function that takes a positive integer n, and returns
 * the sum of the integers from 1 to n.
 *
 *============================================================================*/

func SumOfFirstNIntegers(n int) int {
    return (1 + n) * n / 2
}

/*==============================================================================
 * 2. Speed of a Marathoner
 *
 * Write a function TimeToRun(marathonHours, marathonMinutes, miles) that
 * takes: the time a runner ran a marathon in possibly fractional hours and
 * possibly fractional minutes and a possibly fractional number of miles and
 * return the time in DAYS it should take the runner to run miles if he or she
 * runs at the same pace as they did in the marathon.
 * 
 * For example: TimeToRun(3.1, 23.2, 107.1) should return 0.5938.
 *
 * Your function should also print out the answer in the format:
 *
 *     You could run 107.1 miles in 0.5938 days.
 *
 * Recall that there are 26.2 miles in a marathon.
 *============================================================================*/

// TimeToRun(3.1, 23.2, 107.1) should equal 0.5939
func TimeToRun(marathonHours, marathonMinutes, miles float64) float64 {
    var speed float64 = 26.2 / (marathonHours * 60 + marathonMinutes)
    var days float64 = (miles / speed) / (24 * 60)
    fmt.Println("You could run %.4f in %.4f days.", miles, days)
    return days
}

/*==============================================================================
 * 3. Generalized Fibonacci sequences
 *
 * Implement a function GenFib(a0,a1,n) that takes: two positive integers a0,
 * a1 and a positive integer n, and returns the nth item in the sequence
 * defined by the rule:
 *      a_n = a_{n-1} + a_{n-2}.
 *============================================================================*/

func GenFib(a0, a1, n int) int {
    var a_n int
    for i := 1; i < n; i++ {
        a_n = a0 + a1
        a0 = a1
        a1 = a_n
        }
    return a_n
}

/*==============================================================================
 * 4. Reversing Integers
 *
 * Write a function ReverseInteger(n) that takes an integer, and returns 
 * the integer  formed by reversing the decimal digits of n. For example:
 *      1234 -> 4321
 *      20000 -> 2
 *      1331  -> 1331
 *      -60 -> -6
 *===========================================================================*/

func ReverseInteger(n int) int {
    var sig int = (n >> 31) * 2 + 1
    var n0 int = n * sig
    var res int = 0
    for n0 > 0 {
        var temp = n0 % 10;
        n0 /= 10
        res = res * 10 + temp
    }
    return res * sig
}

/*==============================================================================
 * 5. Growth of a Population
 *
 * The size at time t of a population with a birth rate r can be modeled as:
 *
 *      x_t = r*x_{t-1}(1 - x_{t-1})
 *
 * Write a function PopSize(r, x_0, max_t) that prints out the size of the
 * population (x_t) for t=0...max_t, where x_0 is the initial population size.
 * Assume population size and the birth rate parameter r are real numbers; t is
 * an integer.
 *
 * Your function should also return the final population size.
 *============================================================================*/

func PopSize(r, x_0 float64, max_t int) float64 {
    for t := 0; t < max_t; t++ {
        x_0 = r * x_0 * (1.0 - x_0)
        if x_0 > 1 {
            x_0 = 1
        }
        if x_0 < 0 {
            x_0 = 0
        }
        fmt.Println(x_0)
    }
    return x_0
}

/*==============================================================================
 * 6. Hailstone function
 *
 * The Hailstone function h(n) is defined as n/2 if n is even or 3n+1 if n is
 * odd.  The Hailstone sequence for n is defined by repeatedly applying this
 * function:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * It's conjectured that for all n, this sequence eventually returns to 1.
 * Write a function HailstoneReturnsTo1(n) to compute the smallest number of 
 * times h must be applied to n before the sequence returns to 1.
 *============================================================================*/

func h(n int) int {
    var a int = n % 2
    if a == 1 {
        return n * 3 + 1
    }
    return n / 2
}

func HailstoneReturnsTo1(n int) int {
    var i int = 0
    for n > 1 {
        var temp int = h(n)
        n = temp
        i++
    }
    return i
}

/*==============================================================================
 * 7. Hailstone function maximum
 *
 * Write a function MaxHailstoneValue(n) that takes an integer, and returns the
 * maximum value that the Hailstone sequence:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * achieves before it returns to 1.
 *============================================================================*/

func MaxHailstoneValue(n int) int {
    var max int = n
    for n > 1 {
        var temp int = h(n)
        if temp > max {
            max = temp
        }
        n = temp
    }
    return max
}

/*==============================================================================
 * 8. Find the kth digit of an integer n
 *
 * Implement a function that takes an integer n, and a positive integer k and
 * returns the k-th decimal digit of n, with digit 1 being the rightmost (least
 * significant) digit.
 *
 *============================================================================*/

func KthDigit(n int, k int) int {
    n = n / int(math.Pow(10, float64(k - 1)))
    n = n % 10
    return int(math.Abs(float64(n)))
}

/*===========================================================================
 * 9. Hypergeometric distribution
 *
 * Write a function Hypergeometric(M,N,n,k) that takes 4 integers and returns
 * a float64 which is the value of the hypergeometric distribution 
 *   Pr[red = k] = {M choose k}{N choose n-k} / {M+N choose n}
 *
 * Be careful about overflow: Your funciton should be able to compute
 *      Hypergeometric(5000, 5000, 25, 15)
 *      Hypergeometric(5000, 5000, 50, 15)
 * but not necessarily:
 *      Hypergeometric(5000, 5000, 100, 15)
 *===========================================================================*/

func Hypergeometric(M, N, n, k int) float64 {
    var res float64 = 1
    var fk float64 = float64(k)
    var fn float64 = float64(n)
    var fnk float64 = float64(n - k)
    var i int = M;
    var j int = M + N;
    var s int = N;
    for ; i >= M - k + 1; i-- {
        res = res * float64(i) / fk
        fk--
        //fmt.Println(res)
        for ; j >= M + N - n + 1; j-- {
            res = res / float64(j) * fn
            fn--
            //fmt.Println(res)
            for ; s >= N - n + k + 1; s-- {
                res = res / fnk * float64(s)
                fnk--
                //fmt.Println(res)
            }
        } 
    }
    return res
}