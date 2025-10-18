package main

func productExceptSelf(nums []int) []int {
    pref := make([]int, len(nums))
    suff := make([]int, len(nums))
    pref[0], suff[len(nums)-1] = nums[0], nums[len(nums)-1]
    for i:=1;i<len(nums);i++ {
        pref[i], suff[len(nums)-i-1] = nums[i], nums[len(nums)-i-1]
        pref[i] *= pref[i-1]
        suff[len(nums)-i-1] *= suff[len(nums)-i]
    }
    res := make([]int, len(nums))
    for i:=1;i<len(nums)-1;i++ {
        res[i] = pref[i-1] * suff[i+1]
    }
    res[0] = suff[1]
    res[len(nums)-1] = pref[len(nums)-2]
    return res
}
