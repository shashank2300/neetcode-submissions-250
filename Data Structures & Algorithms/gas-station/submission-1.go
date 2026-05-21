func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
    currentGas := 0
    startIdx := 0

    for i := 0; i < len(gas); i++ {
        net := gas[i] - cost[i]
        
        // Track the overall total to see if a solution is even possible
        totalGas += net
        
        // Track the gas for our current starting point
        currentGas += net

        // If our tank drops below zero, we can't reach the next station
        if currentGas < 0 {
            // Therefore, the current station (and any station before it) 
            // cannot be the start. Move the start to the next station.
            startIdx = i + 1
            
            // RESET the tank for the new starting station!
            currentGas = 0 
        }
    }

    // If total gas is negative, it's impossible to complete the circuit
    if totalGas < 0 {
        return -1
    }

    // Otherwise, the surviving startIdx is mathematically guaranteed to work
    return startIdx
}
