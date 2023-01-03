package n2309

func greatestLetter(s string) string {
	tm := map[int32]int32{}
	var max int32
	for _, c := range s {
		var has int32

		if c <= 'Z' {
			if _, ok := tm[c+32]; ok {
				has = c
			}
		} else {
			if _, ok := tm[c-32]; ok {
				has = c - 32
			}
		}

		if has > max {
			max = has
		}
		tm[c] = 0
	}

	if max > 0 {
		return string(max)
	}
	return ""
}
