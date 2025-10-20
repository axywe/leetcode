package main

func finalValueAfterOperations(operations []string) int {
    result := 0
    for i:=0;i<len(operations);i++ {
        if operations[i] == "X++" || operations[i] == "++X" {
            result++
        } else {
            result--
        }
    }
    return result
}