func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
    total := len(A) + len(B)
    half := (total + 1) / 2

    if len(B) < len(A) {
        A, B = B, A
    }

    l, r := 0, len(A)
    for l <= r {
        i := l + (r-l) / 2
        j := half - i

        Aleft := math.MinInt64
        if i > 0 {
            Aleft = A[i-1]
        }
        Aright := math.MaxInt64
        if i < len(A) {
            Aright = A[i]
        }
        Bleft := math.MinInt64
        if j > 0 {
            Bleft = B[j-1]
        }
        Bright := math.MaxInt64
        if j < len(B) {
            Bright = B[j]
        }

        if Aleft <= Bright && Bleft <= Aright {
            if total%2 != 0 {
                return float64(max(Aleft, Bleft))
            }
            return (float64(max(Aleft, Bleft)) + float64(min(Aright, Bright))) / 2.0
        } else if Aleft > Bright {
            r = i - 1
        } else {
            l = i + 1
        }
    }
    return -1
}
