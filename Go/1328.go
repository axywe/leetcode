package main

func breakPalindrome(palindrome string) string {
    if len(palindrome) == 1 {
        return ""
    }
    newString := ""
    isChanged := false
    for i := 0; i<len(palindrome);i++ {
        sym := palindrome[i]
        if int(sym) == 97 && i == len(palindrome) - 1 && isChanged == false {
            newString += string(int(sym)+1)
            isChanged = true
        } else if isChanged == false && int(sym) != 97 && i != len(palindrome)/2 {
            newString += string(97)
            isChanged = true
        } else {
            newString += string(sym)
        }
    }
    return newString
}