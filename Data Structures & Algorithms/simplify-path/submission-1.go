func simplifyPath(path string) string {
	res := make([]string, 0)
	fmt.Println(strings.Split(path, "/"))
	for _, s := range strings.Split(path, "/") {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(res) > 0 {
				res = res[0:len(res)-1]
			}
		} else {
			res = append(res, s)
		}
		fmt.Println(res)
	}
	if len(res) == 0 {
		return "/"
	}
	return "/"+strings.Join(res, "/")
}
