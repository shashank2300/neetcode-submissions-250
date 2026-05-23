func predictPartyVictory(senate string) string {
    R := []int{}
    D := []int{}
    n := len(senate)

    for i, c := range senate {
        if c == 'R' {
            R = append(R, i)
        } else {
            D = append(D, i)
        }
    }

    for len(R) > 0 && len(D) > 0 {
        rTurn := R[0]
        R = R[1:]
        dTurn := D[0]
        D = D[1:]

        if rTurn < dTurn {
            R = append(R, rTurn+n)
        } else {
            D = append(D, dTurn+n)
        }
    }
	
	if len(R) > 0 {
        return "Radiant"
    }
    return "Dire"
}
