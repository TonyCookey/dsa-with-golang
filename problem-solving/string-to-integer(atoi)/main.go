package string_to_integer_atoi_

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	var str strings.Builder
	for _, v := range s {
		if v >= 48 && v <= 57 {
			str.WriteRune(v)
			continue
		}
		if str.Len() > 0 && (v < 48 || v > 57) {
			break
		}
		if v == ' ' {
			continue
		}
		if str.Len() == 0 && (v == '-' || v == '+') {
			str.WriteRune(v)
			continue
		}

		if v < 48 || v > 57 {
			break
		}
	}
	i, _ := strconv.Atoi(str.String())

	if i > math.MaxInt32 {
		return math.MaxInt32
	}

	if i < math.MinInt32 {
		return math.MinInt32
	}
	return i

}
