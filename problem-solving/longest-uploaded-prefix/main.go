package longest_uploaded_prefix

type LUPrefix struct {
	arr         []bool
	last_prefix int
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		arr:         make([]bool, n+1, n+1),
		last_prefix: 1,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.arr[video] = true
}

func (this *LUPrefix) Longest() int {

	for i := this.last_prefix; i < len(this.arr); i++ {
		if this.arr[i] == false {
			if i > 1 {
				this.last_prefix = i - 1
			}
			return i - 1
		}
	}
	return len(this.arr) - 1
}
