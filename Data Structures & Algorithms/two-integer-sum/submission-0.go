func twoSum(nums []int, target int) []int {
    prev := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        if j, found := prev[diff]; found {
            return []int{j, i}
        }
        prev[n] = i
    }
    return []int{}
}
