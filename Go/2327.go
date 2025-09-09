package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
    known := int64(1)
    arr := make([]int64, n+1)
    arr[1] = 1
    for i:=1;i<n+1;i++{
        if i >= forget {
            known -= arr[i-forget+1]
        }
        if i >= delay {
            j:=i-forget+1
            if j < 0 {
                j = 1
            }
            for ;j<i-delay+1;j++ {
                arr[i] += arr[j]
            }
            known += arr[i]
            arr[i] = arr[i]%1000000007
            known = known%1000000007
        }
    }
    sum := int64(0)
    for i:=n-forget+1;i<n+1;i++ {
        sum += arr[i]
    }
    return int(sum%1000000007)
}