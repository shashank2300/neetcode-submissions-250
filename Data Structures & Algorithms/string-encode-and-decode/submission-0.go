type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	var sizes []string
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")

	var res []string
	i := 0
	for _, size := range sizes {
		if size == "" {
			continue
		}
		length, _ := strconv.Atoi(size)
		res = append(res, parts[1][i:i+length])
		i += length
	}
	return res
}
