package n2042

import (
	"math"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	sa := strings.Split(s, " ")
	ai := math.MinInt
	for _, a := range sa {
		t, err := strconv.Atoi(a)
		if err == nil {
			if t <= ai {
				return false
			}
			ai = t
		}
	}
	return true
}
