package main

func searchMatrix(matrix [][]int, target int) bool {
    var index int
    for l, r := 0, len(matrix) - 1; r >= l; {
        mid := l + (r - l)/2

        if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
            index = mid
            break
        } else if matrix[mid][0] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    for l, r := 0, len(matrix[0]) - 1; r >= l; {
        mid := l + (r - l)/2

        if matrix[index][mid] == target {
            return true
        } else if matrix[index][mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}
