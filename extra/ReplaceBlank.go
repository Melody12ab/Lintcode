package main

import (
    "fmt"
    "errors"
)

//合并两个数组（字符串）时，从前往后需要重复移动数字（字符）多次，可以考虑从后往前移，减少移动次数，提高效率
//length为chars的总容量
func ReplaceBlank(chars [] rune, length int) {
    if chars == nil && length <= 0 {
        return
    }
    originalLength := 0
    numberOfBlank := 0
    for i := 0; i < len(chars); i++ {
        originalLength++
        if chars[i] == ' ' {
            numberOfBlank++
        }
    }
    newLength := originalLength + numberOfBlank*2
    if newLength > cap(chars) {
        errors.New("数组容量不够")
    }
    indexOfOriginal := originalLength - 1
    indexOfNew := newLength - 1
    for indexOfOriginal >= 0 && indexOfNew > indexOfOriginal {
        if chars[indexOfOriginal] == ' ' {
            chars[indexOfNew] = '0'
            indexOfNew--
            chars[indexOfNew] = '2'
            indexOfNew--
            chars[indexOfNew] = '%'
            indexOfNew--
        } else {
            fmt.Println(indexOfNew)
            fmt.Println(indexOfOriginal)
            chars[indexOfNew] = chars[indexOfOriginal]
            indexOfNew--
        }
        indexOfOriginal--
    }

}
