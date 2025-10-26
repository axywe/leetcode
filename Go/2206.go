package main

func divideArray(nums []int) bool {
    m := make(map[int]int)
    unpaired := 0
    for i:=0;i<len(nums);i++ {
        if m[nums[i]] % 2 == 0 {
            unpaired++
            m[nums[i]]++
        } else {
            unpaired--
            m[nums[i]]--
        }
    }
    if unpaired > 0 {
        return false
    } 
    return true
}