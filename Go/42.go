package main

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func trap(height []int) int {
    l := len(height)
    left := make([]int, l)
    right := make([]int, l)
    for i:=1;i<l;i++ {
        left[i] = max(height[i-1], left[i-1])
        right[l-i-1] = max(height[l-i], right[l-i])
    }
    sum := 0
    for i:=0;i<l;i++ {
        cur := min(left[i], right[i]) - height[i]
        if cur > 0 {
            sum += cur
        }
    }
    return sum
}
