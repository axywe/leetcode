package main

func minOperations(nums []int) int {
    same := true
    for i := 1; i < len(nums);i++ {
        if nums[i-1] != nums[i] {
            same = false
        }
    }
    if same == true {
        return 0
    } else {
        return 1
    }
}