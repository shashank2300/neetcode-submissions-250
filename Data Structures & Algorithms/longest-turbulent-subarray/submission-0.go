func maxTurbulenceSize(arr []int) int {
	inc, dec := 1, 1
	ans := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
        	// We just went UP. This extends a previous DOWN trend.
            inc = dec + 1
            dec = 1
    	} else if arr[i] < arr[i-1] {
            // We just went DOWN. This extends a previous UP trend.
            dec = inc + 1
            inc = 1
        } else {
            // Numbers are equal. Turbulence breaks entirely.
            inc = 1
            dec = 1
        }
        
        // The answer is the maximum length seen so far
        ans = max(ans, max(inc, dec))
	}
	return ans
}
