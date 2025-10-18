func calcBox(i int, board [][]byte) bool {
    b := make(map[byte]bool, 9)
    for j:=3*(i/3);j<(3*(i/3)) + 3;j++ {
        for k:=3*(i%3);k<(3*(i%3)) + 3;k++ {
            if board[j][k] != '.' {
                if b[board[j][k]] == true {
                    return false
                }
                b[board[j][k]] = true
            }
        }
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
    for i:=0;i<9;i++ {
        c := make(map[byte]bool, 9)
        r := make(map[byte]bool, 9)
        for j:=0;j<9;j++ {
            if board[i][j] != '.' {
                if r[board[i][j]] == true {
                    return false
                }
                r[board[i][j]] = true
            }
            if board[j][i] != '.' {
                if c[board[j][i]] == true {
                    return false
                }
                c[board[j][i]] = true
            }
        }
        if ! calcBox(i, board) {
            return false
        }

    }
    return true
}
