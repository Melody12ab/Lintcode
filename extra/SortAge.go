package main

import "errors"

//时间复杂度为N的排序
func SortAge(ages []int, length int) {
    if ages == nil && length <= 0 {
        return
    }
    oldestAge := 99
    var timesOfAge [100]int
    for i := 0; i < length; i++ {
        age := ages[i]
        if age < 0 || age > oldestAge {
            errors.New("age out of range")
        }
        timesOfAge[age]++
    }
    index := 0
    for i := 0; i < oldestAge; i++ {
        for j := 0; j < timesOfAge[i]; j++ {
            ages[index] = i
            index++
        }
    }

}
