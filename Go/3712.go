package main

func sumDivisibleByK(nums []int, k int) int {
    m := make(map[int]int)
    count := 0
    for i:=0;i<len(nums);i++ {
        m[nums[i]]++
    }
    for key, v := range m {
        if v % k == 0 {
            count+=key*v
        }
    }
    return count
}