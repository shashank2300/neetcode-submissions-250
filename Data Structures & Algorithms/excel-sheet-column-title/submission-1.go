func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		rem := columnNumber%26
		if rem == 0 {
			rem = 26
		}
		ans = append([]byte{byte('Z'-(26-rem))}, ans...)
		if columnNumber <= 26 {
			break
		}
		columnNumber /= 26
	}

	return string(ans)
}
