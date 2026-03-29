func calPoints(operations []string) int {
	st := make([]int, 0)
	for _, op := range operations {

		if val, err := strconv.Atoi(op); err == nil {
			st = append(st, val)
		} else {
			switch op {
				case "+":
					st = append(st, st[len(st)-2]+st[len(st)-1])
				case "D":
					st = append(st, 2*st[len(st)-1])
				case "C":
					st = st[0:len(st)-1]
			}
		}
	}
	ans := 0
	for _, s := range st {
		ans += s
	}
	return ans
}
