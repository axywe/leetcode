package main

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	maxCost := 27000001
	dist := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dist[i] = append(dist[i], make([]int, 26)...)
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				continue
			}
			dist[i][j] = maxCost
		}
	}

	for i := 0; i < len(original); i++ {
		o := int(original[i] - 'a')
		c := int(changed[i] - 'a')
		if dist[o][c] > cost[i] {
			dist[o][c] = cost[i]
		}
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k] < maxCost && dist[k][j] < maxCost {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}
	sum := int64(0)
	for i := 0; i < len(source); i++ {
		s := int(source[i] - 'a')
		t := int(target[i] - 'a')
		if dist[s][t] == maxCost {
			return -1
		}
		sum += int64(dist[s][t])
	}

	return sum
}
