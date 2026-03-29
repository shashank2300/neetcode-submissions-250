func lastStoneWeight(stones []int) int {
    sort.Ints(stones)
    n := len(stones)

    for n > 1 {
        cur := stones[n-1] - stones[n-2]
        n -= 2

        if cur > 0 {
            pos := sort.Search(n, func(i int) bool {
                return stones[i] >= cur
            })

            for i := n; i > pos; i-- {
                stones[i] = stones[i-1]
            }
            stones[pos] = cur
            n++
        }
    }

    if n > 0 {
        return stones[0]
    }
    return 0
}