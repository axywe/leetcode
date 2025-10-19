package main

func rotate(s string, b int) string {
    slice := []byte(s)
    b = b % len(s)
    first := slice[:len(s)-b]
    second := slice[len(s)-b:]
    return string(append(second, first...))
}


func add(s string, a int) string {
    slice := []byte(s)
    for i:=1;i<len(slice);i+=2 {
        slice[i] = byte((int(slice[i]) - int('0') + a)%10 + int('0'))
    }
    return string(slice)
}

func bfs(s string, a, b int, m map[string]bool) map[string]bool {
    rotated := rotate(s, b)
    if ! m[rotated] {
        m[rotated] = true
        m = bfs(rotated, a, b, m)
    }
    added := add(s, a)
    if ! m[added] {
        m[added] = true
        m = bfs(added, a, b, m)
    }
    return m
}

func lexSmallest(s1, s2 string) string {
    for i := 0; i<len(s1);i++ {
        if s1[i] < s2[i] {
            return s1
        } else if s1[i] > s2[i] {
            return s2
        }
    }
    return s1
}

func findLexSmallestString(s string, a int, b int) string {
    allSeq := make(map[string]bool, 10*10*len(s))
    allSeq[s] = true
    allSeq = bfs(s, a, b, allSeq)
    smallest := s
    for k := range allSeq {
        smallest = lexSmallest(smallest, k)
    }
    return smallest
}