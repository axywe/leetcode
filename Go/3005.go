package main

func maxFrequencyElements(nums []int) int {
    m := make(map[int]int)
    max := 0
    sum := 0
    for _, num := range nums {
        m[num]++
        if m[num] == max {
            sum += m[num]
        }
        if m[num] > max {
            max = m[num]
            sum = m[num]
        }
    }
    return sum
}