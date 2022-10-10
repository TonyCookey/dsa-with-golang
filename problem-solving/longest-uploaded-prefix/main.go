package longest_uploaded_prefix

type LUPrefix struct {
	arr        []bool
	lastPrefix int
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		arr:        make([]bool, n+1, n+1),
		lastPrefix: 1,
	}
}

func (p *LUPrefix) Upload(video int) {
	p.arr[video] = true
}

func (p *LUPrefix) Longest() int {

	for i := p.lastPrefix; i < len(p.arr); i++ {
		if p.arr[i] == false {
			if i > 1 {
				p.lastPrefix = i - 1
			}
			return i - 1
		}
	}
	return len(p.arr) - 1
}
