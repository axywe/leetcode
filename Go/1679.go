package main

func maxOperations(nums []int, k int) int {
    m := make(map[int]int)
    count := 0
    for i:=0;i<len(nums);i++{
        if m[k-nums[i]] > 0 {
            m[k-nums[i]]--
            count++
        } else {
            m[nums[i]]++
        }
    }
    return count
}