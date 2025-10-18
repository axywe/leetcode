package main

func longestConsecutive(nums []int) int {
    m := make(map[int]bool, len(nums))
    for i:=0;i<len(nums);i++ {
        m[nums[i]] = true
    }
    res := 0
    for k := range m {
        if m[k-1] {
            continue
        }

        curr := 0
        for i := k;m[i];i++ {
            curr++
        }

        if curr > res {
            res = curr
        }
    }
    return res
}
