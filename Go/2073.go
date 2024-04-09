package main

func timeRequiredToBuy(tickets []int, k int) int {
	var seconds int = 0
main:
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 {
				tickets[i]--
				seconds++
			}
			if tickets[k] == 0 {
				break main
			}
		}
	}
	return seconds
}
