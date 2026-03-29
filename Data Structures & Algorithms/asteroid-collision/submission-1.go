func asteroidCollision(asteroids []int) []int {
	st := make([]int, 0)


	for _, n := range asteroids {
		if n < 0 {
			put := true
			for len(st) > 0 && st[len(st)-1] > 0 {
				if st[len(st)-1] == -n {
					st = st[:len(st)-1]
					put = false
					break
				}
				if st[len(st)-1] > -n {
					put = false
					break
				}
				if st[len(st)-1] < -n {
					st = st[:len(st)-1]
					put = true
				}
			}
			if put {
				st = append(st, n)
			}
		} else {
			st = append(st, n)
		}
	}
	return st
}
