package sliding_window_maximum

type Deque struct {
	queue []int
}

func (d *Deque) push(val int) {
	d.queue = append(d.queue, val)
}

func (d *Deque) popBack() {
	d.queue = d.queue[:len(d.queue)-1]
}

func (d *Deque) popFront() {
	d.queue = d.queue[1:len(d.queue)]
}

func (d *Deque) getFront() int {
	return d.queue[0]
}

func (d *Deque) getBack() int {
	return d.queue[len(d.queue)-1]
}
func (d *Deque) isEmpty() bool {
	return len(d.queue) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	var left, right int

	deque := &Deque{}

	for right < len(nums) {
		for !deque.isEmpty() && nums[deque.getBack()] < nums[right] {
			deque.popBack()
		}
		deque.push(right)

		if left > deque.getFront() {
			deque.popFront()
		}

		if right+1 >= k {
			result = append(result, nums[deque.getFront()])
			left++
		}
		right++
	}
	return result
}
