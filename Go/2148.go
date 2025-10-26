package main

func countElements(nums []int) int {
    s, l := math.MaxInt32, math.MinInt32
    for i:=0;i<len(nums);i++ {
        if nums[i] < s {
            s = nums[i]
        }
        if nums[i] > l {
            l = nums[i]
        }
    }

    count := 0
    for i:=0;i<len(nums);i++ {
        if nums[i] < l && nums[i] > s {
            count++
        }
    }
    return count
}