package main

func minOperations(nums []int, k int) int {
    c := 0
    for i:=0;i<len(nums);i++{
        if nums[i] < k {
            c++
        }
    }
    return c
}