package main //声明本文件的 package 名
 
import "fmt" //import 语言的 fmt 库——用于输出

func primesSieve(isComposite []bool) {
    var biggestPrime = 2
    for biggestPrime < len(isComposite) {
        for i := 2*biggestPrime; i< len(isComposite); i+=biggestPrime {
            isComposite[i] = true
        }
        biggestPrime++
        for biggestPrime < len(isComposite) && isComposite[biggestPrime] {
            biggestPrime++
        }
    }
}

func main() {
    var composites []bool = make([]bool, 100000000)
    primesSieve(composites)
    var primeCount int = 0
    var primesList []int = make([]int, 0)
    for i, isComp := range composites {
        if !isComp && i >= 2 {
            primeCount++
            fmt.Println("Number of primes <=", i, "is", primeCount)
            primesList = append(primesList, i)
        }
    }
}