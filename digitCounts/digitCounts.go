/**
 * @param k: An integer
 * @param n: An integer
 * @return: An integer denote the count of digit k in 1..n
 */
import(
    "strings"
    "strconv"
)

func digitCounts (k int, n int) int {
    // write your code here
    sum:=0
    kstr:=strconv.Itoa(k)
    for i:=0;i<=n;i++{
        sum+=strings.Count(strconv.Itoa(i),kstr)
    }
    return sum
}

