package main

func search(nums []int, target int) int {
    for i, j := 0, len(nums)-1; i <= j; {
        mid := (j+i)/2
        if nums[mid] == target {
            return mid
        } 
        if nums[mid] > target {
            j = mid - 1
        } else {
            i = mid + 1
        }
    }
    return -1
}
