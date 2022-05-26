package implement_stacks_using_queues

type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.Queue1 = append(this.Queue1, x)
}

func (this *MyStack) Pop() int {
	leng := len(this.Queue1) - 1
	for i := 0; i < leng; i++ {
		this.Queue2 = append(this.Queue2, this.Queue1[0])
		this.Queue1 = this.Queue1[1:]
	}
	top := this.Queue1[0]
	this.Queue1 = this.Queue2
	this.Queue2 = []int{}
	return top
}

func (this *MyStack) Top() int {
	top := this.Pop()
	this.Push(top)
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.Queue1) == 0
}
