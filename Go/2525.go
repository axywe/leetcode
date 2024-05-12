package main

func categorizeBox(length int, width int, height int, mass int) string {
	bulky, heavy := false, false
	if (length >= 10000 || width >= 10000 || height >= 10000) || int64(length*height*width) >= 1000000000 {
		bulky = true
	}
	if mass >= 100 {
		heavy = true
	}

	if bulky && heavy {
		return "Both"
	} else if !bulky && !heavy {
		return "Neither"
	} else if bulky && !heavy {
		return "Bulky"
	} else {
		return "Heavy"
	}
}
