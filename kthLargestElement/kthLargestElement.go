package main

/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 */
func kthLargestElement(n int, nums []int) int {
    return findkthLargestElement(n, nums, 0, len(nums)-1)
}

func findkthLargestElement(k int, nums []int, begin, end int) int {
    if begin == end {
        return nums[begin]
    }
    index := partition(nums, begin, end+1)
    //end-index+1=n 为num中第n大的数
    if end-index+1 == k {
        return nums[index]
    } else if (end-index+1 > k) {
        return findkthLargestElement(k, nums, index+1, end)
    } else {
        return findkthLargestElement(k-(end-index+1), nums, begin, index-1)
    }

}

//begin为初始索引，end为数组元素个数
func partition(nums []int, begin, end int) int {
    pvalue := nums[begin]
    i := begin
    j := begin + 1
    for j < end {
        if nums[j] < pvalue {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
        j++
    }
    nums[i], nums[begin] = nums[begin], nums[i]
    return i
}
