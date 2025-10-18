package main

func twoSum(numbers []int, target int) []int {
    for i, j := 0, len(numbers)-1;i<j; {
        if numbers[i] + numbers[j] == target {
            return []int{i+1, j+1}
        }
        if numbers[i+1] + numbers[j] <= target {
            i++
        } else {
            j--
        }
    }
    return []int{}
}
