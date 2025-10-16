package main

func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    for i := 0; i < len(strs); i++ {
        arr := [26]int{}
        for j:=0;j<len(strs[i]);j++ {
            arr[strs[i][j] - 'a']++
        }
        m[arr] = append(m[arr], strs[i])
    }
    
    var result [][]string
    for _, g := range m {
        result = append(result, g)
    }
    return result
}
