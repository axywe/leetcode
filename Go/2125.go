package main

func numberOfBeams(bank []string) int {
    count := 0
    latest := 0
    for i:=0;i<len(bank);i++ {
        cur := 0
        for j:=0;j<len(bank[i]);j++{
            if bank[i][j] == '1' {
                count += latest
                cur++
            }
        }
        if cur != 0 {
            latest = cur
        }
    }
    return count
}