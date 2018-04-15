package main

import (
    "fmt"
)

//快排
func partition(nums []int, start, end int) int {
    pivot := nums[start]
    front := start
    behind := start+1
    for behind <= end {
        if nums[behind] < pivot {
            front++
            nums[front], nums[behind] = nums[behind], nums[front]
        }
        behind++
    }
    nums[front], nums[start] = nums[start], nums[front]
    return front
}

func QuickSort(nums []int, start, end int) {
    if start < end {
        mid := partition(nums, start, end)
        QuickSort(nums, start, mid-1)
        QuickSort(nums, mid+1, end)
    }
}

//todo 归并排序

func main() {
    nums := []int{7, 1, 8, 3, 6, 12, 0, 17, 10}
    fmt.Println(nums)
    QuickSort(nums, 0, len(nums)-1)
    fmt.Println(nums)

}
