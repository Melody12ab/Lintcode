package main

import "errors"

//旋转数组的最小数字
//二分查找的思想


func Min(nums []int, length int) int {
    if nums == nil || length < 0 {
        errors.New("Invalid parameters")
    }
    index1 := 0
    index2 := length - 1
    indexMid := index1
    for nums[index1] >= nums[index2] {
        if index2-index1 == 1 {
            indexMid = index2
            break
        }

        indexMid = (index1 + index2) / 2

        //如果下表为index1、index2和indexMid的三个数字相等，则只能顺序查找

        if nums[indexMid] >= nums[index1] {
            index1 = indexMid
        } else if nums[indexMid] <= nums[index2] {
            index2 = indexMid
        }
    }
    return nums[indexMid]
}
