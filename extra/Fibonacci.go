package main

//求第N个Fibonacci数列

//低效的解法，递归中存在大量重复计算
func FibonacciLow(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return FibonacciLow(n-1) + FibonacciLow(n-2)
}

func FibonacciHigh(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    fibNMinusOne := 1
    fibNMinusTwo := 0
    fibN := 0
    for i := 2; i < n; i++ {
        fibN = fibNMinusOne + fibNMinusTwo
        fibNMinusTwo = fibNMinusOne
        fibNMinusOne = fibN
    }
    return fibN

}
