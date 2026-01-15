package main

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	curH, maxCurH := 1, 1
	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			curH++
			if curH > maxCurH {
				maxCurH = curH
			}
		} else {
			curH = 1
		}
	}

	curV, maxCurV := 1, 1
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			curV++
			if curV > maxCurV {
				maxCurV = curV
			}
		} else {
			curV = 1
		}
	}

	return min(maxCurV+1, maxCurH+1) * min(maxCurV+1, maxCurH+1)
}
