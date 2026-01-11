package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := []int{0}
	max := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for j := len(stack) - 1; j >= 0; j-- {
				if heights[i] <= heights[stack[j]] {
					top := heights[stack[len(stack)-1]]
					stack = stack[:len(stack)-1]
					if len(stack) != 0 {
						rect := top * (i - stack[len(stack)-1] - 1)
						if max < rect {
							max = rect
						}
					} else if max < top*i {
						max = top * i
					}
				} else {
					break
				}
			}
			stack = append(stack, i)
		}
	}
	return max
}
