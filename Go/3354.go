package main

func algorithm(nums []int, curr int, dir int) int {
    // dir == -1 -> left
    // dir == 1 -> right
    for ;; {
        curr += dir
        if curr > len(nums)-1 || curr < 0 {
            break
        }

        if nums[curr] > 0 {
            nums[curr]--
            dir *= -1
        }
    }
    for i:=0;i<len(nums);i++ {
        if nums[i] != 0 {
            return 0
        }
    }
    return 1
}

func countValidSelections(nums []int) int {
    count := 0
    cp := make([]int, len(nums))
    for i:=0;i<len(nums);i++ {
        if nums[i] == 0 {
	        copy(cp, nums)
            count += algorithm(cp, i, -1)
            copy(cp, nums)
            count += algorithm(cp, i, 1)
        }
    }
    return count
}