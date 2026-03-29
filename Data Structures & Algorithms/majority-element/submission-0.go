func majorityElement(nums []int) int {
    n := len(nums)
    m := make(map[int]int)

    for _, num := range nums {
        m[num]++
        if m[num] >= 1+n/2 {
            return num
        }
    }

    return -1
}
