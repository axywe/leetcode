package main

func distanceTraveled(mainTank int, additionalTank int) int {
	distance := 0
	count := 0
	for mainTank > 0 {
		mainTank--
		count++
		distance += 10
		if count == 5 && additionalTank >= 1 {
			mainTank++
			additionalTank--
			count = 0
		}
	}
	return distance
}
