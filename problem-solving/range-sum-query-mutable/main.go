package range_sum_query_mutable

type NumArray struct {
	nums []int
	memo map[[2]int]int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
		memo: make(map[[2]int]int),
	}
}

func (this *NumArray) Update(index int, val int) {

	diff := val - this.nums[index]
	this.nums[index] = val
	for i, _ := range this.memo {
		if index >= i[0] && index <= i[1] {
			this.memo[i] += diff
		}
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if val, exists := this.memo[[2]int{left, right}]; exists {
		return val
	}
	var sum int
	for _, v := range this.nums[left : right+1] {
		sum += v
	}
	this.memo[[2]int{left, right}] = sum
	return sum
}
