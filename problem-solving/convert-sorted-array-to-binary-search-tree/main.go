package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)

}
func sortedArrayToBSTHelper(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low)/2

	curr := &TreeNode{
		Val: nums[mid],
	}

	curr.Left = sortedArrayToBSTHelper(nums, low, mid-1)
	curr.Right = sortedArrayToBSTHelper(nums, mid+1, high)
	return curr
}
