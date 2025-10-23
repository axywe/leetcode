package main

func hasSameDigits(s string) bool {
    for ;len(s)>2; {
        str := []byte{}
        for i:=0;i<len(s)-1;i++ {
            current := (int(s[i] - '0') + int(s[i+1] - '0')) % 10 + int('0')
            str = append(str, byte(current))
        }
        s = string(str)
    }
    if s[0] == s[1] {
        return true
    }
    return false
}