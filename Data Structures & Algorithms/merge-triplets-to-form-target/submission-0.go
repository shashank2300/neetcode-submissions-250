func mergeTriplets(triplets [][]int, target []int) bool {
	d := map[int]bool{0:false, 1:false, 2:false}
	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}
		if triplet[0] == target[0] {
			d[0] = true
		}
		if triplet[1] == target[1] {
			d[1] = true
		}
		if triplet[2] == target[2] {
			d[2] = true
		}

		if d[0] && d[1] && d[2] {
			return true
		}
	}
	return false
}
