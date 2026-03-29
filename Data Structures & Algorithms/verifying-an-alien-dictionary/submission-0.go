func isAlienSorted(words []string, order string) bool {
    orderIndex := make([]int, 26)
    for i, c := range order {
        orderIndex[c-'a'] = i
    }

    for i := 0; i < len(words)-1; i++ {
        w1, w2 := words[i], words[i+1]

        for j := 0; j < len(w1); j++ {
            if j == len(w2) {
                return false
            }

            if w1[j] != w2[j] {
                if orderIndex[w1[j]-'a'] > orderIndex[w2[j]-'a'] {
                    return false
                }
                break
            }
        }
    }
    return true
}