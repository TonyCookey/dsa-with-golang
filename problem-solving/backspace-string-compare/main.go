package backspace_string_compare

import "strings"

func backspaceCompare(s string, t string) bool {
	var sx, tx strings.Builder
	var temp string

	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			sx.WriteByte(s[i])
		} else {
			temp = sx.String()
			sx.Reset()
			if temp != "" {
				sx.WriteString(temp[:len(temp)-1])
			}
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] != '#' {
			tx.WriteByte(t[i])
		} else {
			temp = tx.String()
			tx.Reset()
			if temp != "" {
				tx.WriteString(temp[:len(temp)-1])
			}
		}
	}

	return sx.String() == tx.String()
}
