package main

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    m := make(map[int][]int)
    for i:=0;i<len(nums);i++ {
        m[nums[i]] = append(m[nums[i]], i)
    }

    resultMap := make(map[[3]int]bool)
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++ {
            for k:=0;k<3 && k<len(m[-(nums[i]+nums[j])]);k++ {
                if m[-(nums[i]+nums[j])][k] != i && m[-(nums[i]+nums[j])][k] != j {
                    batch := []int{nums[i], nums[j], -(nums[i]+nums[j])}
                    sort.Ints(batch)
                    resultMap[[3]int{batch[0], batch[1], batch[2]}] = true
                }
            }
        }
    }
    var result [][]int
    for k, _ := range resultMap {
        result = append(result, []int{k[0], k[1], k[2]})
    }
    return result
}
