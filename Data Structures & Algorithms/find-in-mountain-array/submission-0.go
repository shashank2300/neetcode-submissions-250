/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    length := mountainArr.length()

	l, r, peak := 1, length-2, 0
	for l <= r {
		m := l + (r-l)/2
		left := mountainArr.get(m-1)
		mid := mountainArr.get(m)
		right := mountainArr.get(m+1)

		if left < mid && mid < right {
			l = m + 1
		} else if left > mid && mid > right {
			r = m-1
		} else {
			peak = m
			break
		}
	}

	// left part
	l, r = 0, peak
	for l <= r {
		m := l + (r-l)/2
		val := mountainArr.get(m)
		if val < target {
			l = m+1
		} else if val > target {
			r = m-1
		} else {
			return m
		}
	}

	// right part
	l, r = peak+1, length-1
	for l <= r {
		m := l + (r-l)/2
		val := mountainArr.get(m)
		if val > target {
			l = m+1
		} else if val < target {
			r = m-1
		} else {
			return m
		}
	}

	return -1
}
