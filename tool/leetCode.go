package tool

import (
	"fmt"
	"math"
)

func CountGoodRectangles(rectangles [][]int) int {
	lenMap := make(map[int]int)
	maxLen := 0
	for i := 0; i < len(rectangles); i++ {
		minL := min(rectangles[i][0], rectangles[i][1])
		maxLen = max(maxLen, minL)
		lenMap[minL]++
	}
	res := lenMap[maxLen]
	fmt.Println(res)
	return res
}

func constructRectangle(area int) []int {
	l := int(math.Floor(math.Sqrt(float64(area))))
	for i := l; i < area; i++ {
		w := area / i
		if w*i == area {
			return []int{i, w}
		}
	}
	return nil
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftH := Dfs(root.left)
	rightH := Dfs(root.right)
	if math.Abs(float64(leftH)-float64(rightH)) > 1 {
		return false
	}
	return isBalanced(root.left) && isBalanced(root.right)
}

func Dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := Dfs(root.left)
	right := Dfs(root.right)
	return max(left, right) + 1
}

func RemoveDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	res := make([]int, size)
	index := 0
	res[0] = nums[0]
	for i := 1; i < size; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		res[index] = nums[i]
		index++
	}

	for i := 0; i < index; i++ {
		nums[i] = res[i]
	}
	return index
}

func main() {

}
