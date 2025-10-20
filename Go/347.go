package main

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        m[nums[i]]++
    }
    c := make([][]int, len(nums)+1)
    for k, v := range m {
        c[v] = append(c[v], k)
    }

    res := []int{}

    for i:=len(c)-1;i>0;i-- {
        for j := 0;j<len(c[i]) && k>0;j++ {
            res = append(res, c[i][j])
            k--
        }
    }
    return res
}
