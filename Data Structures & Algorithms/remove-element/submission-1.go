func removeElement(nums []int, val int) int {
    sort.Ints(nums)
    n := len(nums)
    rem := 0
    for i := range nums {
        if n-1-rem < i {
            break
        }
        if nums[i] == val {
            if nums[i] == nums[n-1-rem] {
                rem += (n-1-rem-i+1)
                break
            }
            nums[i], nums[n-1-rem] = nums[n-1-rem], nums[i]
            rem++
        }
    }
    return n-rem
}
