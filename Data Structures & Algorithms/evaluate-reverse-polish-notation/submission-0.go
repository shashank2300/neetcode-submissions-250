func evalRPN(tokens []string) int {
	st := make([]int, 0)

	for _, c := range tokens {
		d, err := strconv.Atoi(c)
		switch {
			case err == nil:
				st = append(st, d)
			case c == "+":
				t := st[len(st)-1] + st[len(st)-2]
				st[len(st)-2] = t
				st = st[:len(st)-1]
			case c == "*":
				t := st[len(st)-1] * st[len(st)-2]
				st[len(st)-2] = t
				st = st[:len(st)-1]
			case c == "-":
				t := st[len(st)-2] - st[len(st)-1]
				st[len(st)-2] = t
				st = st[:len(st)-1]
			case c == "/":
				t := st[len(st)-2] / st[len(st)-1]
				st[len(st)-2] = t
				st = st[:len(st)-1]
		}
	}
	return st[0]
}
