/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func aplusb (a int, b int) int {
    and := a&b  //进位
    xor := a^b  //不算进位的按位和
    for and != 0 {
        bitsum := xor
        bitwei := and<<1
        and = bitsum & bitwei
        xor = bitsum ^ bitwei
    }
    return xor
}

