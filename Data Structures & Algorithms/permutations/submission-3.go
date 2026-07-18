import "slices"

func permute(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{{}}
    }

    res := make([][]int, 0)
    perms := permute(nums[1:])
    for _, perm := range perms {
        for i := 0; i<=len(perm); i++ {
            pCopy := slices.Clone(perm)
            pCopy = slices.Insert(pCopy, i, nums[0])
            res = append(res, pCopy)
        }
    }

    return res
}
