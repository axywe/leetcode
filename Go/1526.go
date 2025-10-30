package main

func minNumberOperations(target []int) int {
    count := 0
    for i:=1;i<len(target);i++ {
        if target[i] < target[i-1] {
            count+=target[i-1]-target[i]
        }
    }
    return count + target[len(target)-1]
}