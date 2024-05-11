package main

func thirdMax(nums []int) int {
	f, s, t := -2147483649, -2147483649, -2147483649
	for _, i := range nums {
		if i != f && i != s && i != t {
			if i > f {
				t, s, f = s, f, i
			} else if i > s {
				t, s = s, i
			} else if i > t {
				t = i
			}
		}
	}
	if t == -2147483649 {
		return f
	} else {
		return t
	}
}
