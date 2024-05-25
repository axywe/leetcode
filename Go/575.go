package main

func distributeCandies(candyType []int) int {
	n := len(candyType) / 2
	m := make(map[int]int)
	for i := 0; i < n*2; i++ {
		m[candyType[i]]++
	}
	if len(m) > n {
		return n
	}
	return len(m)
}
