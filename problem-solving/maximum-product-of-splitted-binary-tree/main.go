package maximum_product_of_splitted_binary_tree

import "math"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func maxProduct(root *TreeNode) int {
	var totalSum, result int

	var getSum func(root *TreeNode) int

	getSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		subTreeSum := root.Val + getSum(root.Left) + getSum(root.Right)

		result = max(result, (totalSum-subTreeSum)*subTreeSum)
		return subTreeSum
	}

	totalSum = getSum(root)
	getSum(root)

	return result % int(math.Pow10(9)+7)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
