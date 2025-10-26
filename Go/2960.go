package main

func countTestedDevices(batteryPercentages []int) int {
    count := 0
    for i:=0;i<len(batteryPercentages);i++ {
        if batteryPercentages[i] - count > 0 {
            count++
        }
    }
    return count
}