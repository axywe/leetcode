package main

func maxDistinctElements(nums []int, k int) int {
    m := make(map[int]bool)
    sort.Ints(nums)
    count := 0
    min := nums[0]-k
    for i:=0;i<len(nums);i++ {
        if min < nums[i]-k {
            min = nums[i]-k
        }
        for j:=min;j<=nums[i]+k;j++ {
            if ! m[j] {
                m[j] = true
                min = j
                count++
                break
            }
        }
    }
    return count
}