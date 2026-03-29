func isValid(s string) bool {
    st := make([]rune, 0)
	for _, c := range s {
		switch  {
			case c=='[' || c=='{' || c=='(':
				st = append(st, c)
			case c==']':
				if len(st) == 0 || st[len(st)-1] != '[' {
					return false
				} else {
					st = st[0:len(st)-1]
				}
			case c=='}':
				if len(st)==0 || st[len(st)-1] != '{' {
					return false
				} else {
					st = st[0:len(st)-1]
				}
			case c==')':
				if len(st)==0 || st[len(st)-1] != '(' {
					return false
				} else {
					st = st[0:len(st)-1]
				}
		}
	}
	if len(st) > 0 {
		return false
	}
	return true
}
