/**
 * @param n: A long integer
 * @return: An integer, denote the number of trailing zeros in n!
 */
func trailingZeros (n int64) int64 {
    // write your code here, try to do it without arithmetic operators.
    var n5  int64 = 5
    var sum int64 = 0
    for n >= n5{
        sum += n/n5
        n5 *= 5
    }
    return sum
}
