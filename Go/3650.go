package main

import "math"

func minCost(n int, edges [][]int) int {
	m := make([][][]int, n)
	for i := 0; i < len(edges); i++ {
		m[edges[i][0]] = append(m[edges[i][0]], []int{edges[i][1], edges[i][2]})
		m[edges[i][1]] = append(m[edges[i][1]], []int{edges[i][0], 2 * edges[i][2]})
	}
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 1; i < n; i++ {
		dist[i] = math.MaxInt
	}
	buckets := make([][]int, 2001)
	buckets[0] = append(buckets[0], 0)
	count := 1
	currentDist := 0
	cur := 0
	for count != 0 {
		for ; len(buckets[(currentDist)%2001]) == 0; currentDist++ {
		}
		cur = buckets[currentDist%2001][len(buckets[currentDist%2001])-1]
		buckets[currentDist%2001] = buckets[currentDist%2001][:len(buckets[currentDist%2001])-1]
		count--

		for i := 0; i < len(m[cur]); i++ {
			if dist[m[cur][i][0]] > dist[cur]+m[cur][i][1] {
				dist[m[cur][i][0]] = dist[cur] + m[cur][i][1]
				if !visited[m[cur][i][0]] {
					newDist := dist[m[cur][i][0]]
					buckets[newDist%2001] = append(buckets[newDist%2001], m[cur][i][0])
					count++
				}
			}
		}
		visited[cur] = true
	}
	if dist[n-1] == math.MaxInt {
		return -1
	}
	return dist[n-1]
}
