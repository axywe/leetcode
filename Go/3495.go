package main

func minOperations(queries [][]int) int64 {
    result := int64(0)
    powers := make([]int64, 16)
    powers[0] = 1
    for i := 1; i < 16; i++ {
        powers[i] = powers[i-1] * 4
    }
    for i := 0; i < len(queries); i++ {
        queryResult := int64(0)
        l := int64(queries[i][0])
        r := int64(queries[i][1])
        for j := 1; j < 16; j++ {
            L := powers[j-1]
            R := powers[j] - 1
            
            left := l
            if L > left {
                left = L
            } 
            right := r
            if R < right {
                right = R
            }

            if right >= left {
                queryResult += (int64(right) - int64(left) + int64(1)) * int64(j)
            } else if L > r {
                break
            }
        }
        if queryResult % 2 == 1 {
            queryResult += 1
        }
        result += queryResult / 2
    }
    return result
}