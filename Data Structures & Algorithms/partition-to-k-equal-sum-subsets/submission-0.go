func canPartitionKSubsets(nums []int, k int) bool {
    total := 0
    n := len(nums)
    for _, num := range nums {
        total += num
    }
    if total % k != 0 {
        return false
    }
    target := total/k
    buckets := make([]int, k)

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    fmt.Println(nums)
    var dfs func(idx int) bool
    dfs = func(idx int) bool {
        if idx >= n {
            return true
        }

        for i := range k {
            if buckets[i] + nums[idx] <= target {
                buckets[i] += nums[idx]
                if dfs(idx+1) {
                    return true
                }
                buckets[i] -= nums[idx]
            }
        }
        return false
    }

    return dfs(0)
}
