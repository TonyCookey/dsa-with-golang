package range_sum_query_immutable

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}
