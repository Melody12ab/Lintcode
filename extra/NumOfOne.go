package main

//计算给定整数N二进制表示中1的个数

//此解法没有给出n为负时的处理方式
func NumOfOneLow(n int) int {
    count := 0
    for n > 0 {
        if n&1 == 1 {
            count++
        }
        n = n >> 1
    }
    return count
}

//下面的解法通过将1每次左移来避免了负数右移可能出现的死循环
//不过改解法中循环的次数等于二进制的位数，32位就是32次。。。。
func NumOfOneHigh(n int) int {
    count := 0
    flag := 1
    for flag == 1 {
        if n&flag == 1 {
            count++
        }
        flag = flag << 1
    }
    return count
}

//一个整数减去1再和原数做位与，会把整数最右边一个1编程0
//todo 很多二进制问题都可以用这个思路解
func NumOfOneHighest(n int) int {
    count := 0
    for n > 0 {
        n = (n - 1) & n
        count++
    }
    return count
}
