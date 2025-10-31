package main

func getSneakyNumbers(nums []int) []int {
    m := make(map[int]int)
    res := []int{}
    for i:=0;i<len(nums);i++ {
        if m[nums[i]] == 1 {
            res = append(res, nums[i])
        }
        m[nums[i]]++
    }
    return res
}