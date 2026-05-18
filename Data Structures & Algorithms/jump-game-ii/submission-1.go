func jump(nums []int) int {
	if len(nums) <= 1 {
        return 0
    }

    jumps := 0
    currentEnd := 0  // The end of the current jump window
    farthest := 0    // The farthest index we can reach from the current window

    // Notice we stop at len(nums)-1. We don't need to jump FROM the last element.
    for i := 0; i < len(nums)-1; i++ {
        // As we move through the current window, keep updating the farthest reach
        farthest = max(farthest, i+nums[i])

        // When we hit the end of the current window, we MUST take a jump
        if i == currentEnd {
            jumps++
            currentEnd = farthest // Set the boundary for the next jump window
            
            // Minor optimization: if our new window reaches the end, we can stop early
            if currentEnd >= len(nums)-1 {
                break
            }
        }
    }

    return jumps
}
