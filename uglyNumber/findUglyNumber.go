package main

func Min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func FindUgly(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1
	index2 := 0
	index3 := 0
	index5 := 0
	for i := 1; i < n; i++ {
		ugly[i] = Min(ugly[index2]*2, ugly[index3]*3, ugly[index5]*5)
		if ugly[i] == ugly[index2]*2 {
			index2++
		}
		if ugly[i] == ugly[index3]*3 {
			index3++
		}
		if ugly[i] == ugly[index5]*5 {
			index5++
		}
	}
	return ugly[n-1]
}

func main() {
	println(FindUgly(6))
	println(FindUgly(7))
}
