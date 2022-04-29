package add_to_array_form_of_integer

func addToArrayForm(num []int, k int) []int {
	var sum, carry int
	var result []int
	i := len(num) - 1
	for k != 0 || i > -1 {
		if i > -1 {
			sum += num[i]
		}
		sum += carry + (k % 10)

		result = append(result, sum%10)
		carry = sum / 10
		sum = 0
		k /= 10
		i--
	}
	if carry == 1 {
		result = append(result, 1)
	}
	return reverse(result)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
