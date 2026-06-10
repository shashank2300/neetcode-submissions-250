func leastInterval(tasks []byte, n int) int {
    counts := make([]int, 26)
    for _, task := range tasks {
        counts[task-'A']++
    }

    sort.Ints(counts)

    maxFreq := counts[25]
    
    // Count how many tasks share this maximum frequency
    maxFreqCount := 1
    for i := 24; i >= 0; i-- {
        if counts[i] == maxFreq {
            maxFreqCount++
        } else {
            break
        }
    }

    // Calculate minimum time using the formula
    ans := (maxFreq - 1) * (n + 1) + maxFreqCount

    // Return the maximum of the calculated time or the total number of tasks
    if ans < len(tasks) {
        return len(tasks)
    }
    
    return ans
}
