package main

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	for i := 0; i < n; i++ {
		for j := 0; (n%2 == 0 && j < n/2) || (n%2 == 1 && j < n/2+1); j++ {
			image[i][j], image[i][n-j-1] = image[i][n-j-1], image[i][j]
			image[i][j], image[i][n-j-1] = (image[i][j]+1)%2, (image[i][n-j-1]+1)%2
		}
	}
	return image
}
