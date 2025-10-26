package main

func secondHighest(s string) int {
    first, second := -1, -1;
    for i := 0;i<len(s);i++ {
        if s[i] >= '0' && s[i] <= '9' {
            if int(s[i] - '0') > first {
                second = first
                first = int(s[i] - '0')
            } else if int(s[i] - '0') > second && int(s[i] - '0') != first {
                second = int(s[i] - '0')
            }
        }
    }
    return second
}