package min_stack

import "fmt"

type MinStack struct {
	stack   []int
	minVals []int
}

func Constructor() MinStack {
	return MinStack{
		stack:   []int{},
		minVals: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.minVals = append(this.minVals, val)
		this.stack = append(this.stack, val)

	} else {
		if val <= this.GetMin() {
			this.minVals = append(this.minVals, val)
		}
		this.stack = append(this.stack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		if this.Top() == this.GetMin() && len(this.minVals) > 0 {
			this.minVals = this.minVals[:len(this.minVals)-1]
		}
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	fmt.Println(this.minVals)
	if len(this.minVals) == 0 {
		return 0
	}
	return this.minVals[len(this.minVals)-1]
}
