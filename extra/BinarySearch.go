package main

import (
    "fmt"
)

func main() {
    lookingFor := 6
    sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
    index := binarySearch(sortedList, lookingFor)
    if index >= 0 {
        fmt.Println("Find the Data:", index)
    } else {
        fmt.Println("Not Find the Data!")
    }
}
func binarySearch(sortedList []int, lookingFor int) int {
    lo := 0
    hi := len(sortedList) - 1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        midValue := sortedList[mid]
        if midValue == lookingFor {
            return midValue
        } else if midValue > lookingFor {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return -1
}
