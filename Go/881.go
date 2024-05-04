package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ptr1, ptr2 := 0, len(people)-1
	fmt.Println(people, ptr1, ptr2)
	boats := 0
	for ptr1 <= ptr2 {
		if people[ptr2]+people[ptr1] <= limit {
			ptr2--
			ptr1++
			boats++
		} else {
			boats++
			ptr2--
		}
	}
	return boats
}
