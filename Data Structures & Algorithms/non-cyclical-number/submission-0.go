func isHappy(n int) bool {
	vis := make(map[int]bool)
	var helper func(n int, vis map[int]bool) bool
	helper = func(n int, vis map[int]bool) bool {
		if vis[n] {
			return false
		}
		if n == 1 {
			return true
		}
		vis[n] = true
		num := 0
		for n > 0 {
			cur := n % 10
			num += cur*cur
			n /= 10
		}
		return helper(num, vis)
	}
	return helper(n, vis)
}
