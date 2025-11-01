package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] < 0 {
			d := false
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] < 0 {
					break
				}
				if -asteroids[i] > stack[j] {
					stack = stack[:j]
				} else if -asteroids[i] == stack[j] {
					stack = stack[:j]
					d = true
					break
				} else {
					d = true
					break
				}
			}
			if !d {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}
