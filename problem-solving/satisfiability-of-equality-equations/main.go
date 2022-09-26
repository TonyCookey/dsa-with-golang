package satisfiability_of_equality_equations

func equationsPossible(equations []string) bool {
	var list uf
	for i, _ := range list {
		list[i] = byte('a' + i)
	}

	for _, eq := range equations {
		if eq[1] == '=' {
			a := list.find(eq[0])
			b := list.find(eq[3])
			if a != b {
				list.union(a, b)
			}
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			if list.find(eq[0]) == list.find(eq[3]) {
				return false
			}
		}
	}
	return true
}

type uf [26]byte

func (list *uf) find(letter byte) byte {
	return list[int(letter-'a')]
}

func (list *uf) union(l1 byte, l2 byte) {
	if l1 > l2 {
		list.union(l2, l1)
		return
	}
	for i, v := range list {
		if v == l2 {
			list[i] = l1
		}
	}
}
