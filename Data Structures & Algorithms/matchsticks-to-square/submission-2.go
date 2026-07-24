func makesquare(matchsticks []int) bool {
    total := 0
    for _, t := range matchsticks {
        total += t
    }
    
    if total == 0 || total%4 != 0 {
        return false
    }
    
    target := total / 4
    
    // Optimization 1: Sort in descending order.
    // Trying to place the largest matchsticks first helps us fail faster 
    // if a solution is impossible, drastically reducing the search space.
    sort.Slice(matchsticks, func(i, j int) bool {
        return matchsticks[i] > matchsticks[j]
    })
    
    // sides represents the current length of our 4 square sides
    sides := make([]int, 4)
    
    var dfs func(index int) bool
    dfs = func(index int) bool {
        // Base case: If we've successfully placed all matchsticks
        if index == len(matchsticks) {
            return true
        }
        
        // Try placing the current matchstick into each of the 4 sides
        for i := 0; i < 4; i++ {
            // Optimization 2: Skip identical side lengths to avoid redundant work.
            // If putting the matchstick in sides[i-1] failed, and sides[i] has 
            // the same current length, putting it in sides[i] will also fail.
            if i > 0 && sides[i] == sides[i-1] {
                continue
            }
            
            // If the matchstick fits in the current side
            if sides[i] + matchsticks[index] <= target {
                sides[i] += matchsticks[index] // Place it
                
                if dfs(index + 1) { // Recurse to the next matchstick
                    return true
                }
                
                sides[i] -= matchsticks[index] // Backtrack if it didn't lead to a solution
            }
        }
        
        return false
    }
    
    return dfs(0)
}