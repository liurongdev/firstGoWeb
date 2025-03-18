package tool

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func DeleteNode(root *TreeNode, key int) *TreeNode {
	return deleteWithParent(root, nil, key)

}

func deleteWithParent(root *TreeNode, parent *TreeNode, key int) *TreeNode {
	target, parent := findNode(root, parent, key)
	if target == nil {
		return root
	}
	if target.Left == nil && target.Right == nil {
		if parent == nil {
			return nil
		}
		if parent.Left == target {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if target.Left == nil && target.Right != nil {
		if parent == nil {
			return target.Right
		}
		if parent.Left == target {
			parent.Left = target.Right
		} else {
			parent.Right = target.Right
		}
	} else if target.Right == nil && target.Left != nil {
		if parent == nil {
			return target.Left
		}
		if parent.Left == target {
			parent.Left = target.Left
		} else {
			parent.Right = target.Left
		}
	} else {
		nextNode := findNext(target)
		if nextNode != nil {
			target.Val = nextNode.Val
			deleteWithParent(target.Right, target, nextNode.Val)
		}
	}
	return root
}

func findNode(root *TreeNode, parent *TreeNode, key int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Val == key {
		return root, parent
	} else if key < root.Val {
		return findNode(root.Left, root, key)
	} else {
		return findNode(root.Right, root, key)
	}
	return nil, nil
}

func findNext(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return nil
	}
	var next = root.Right
	for next.Left != nil {
		next = next.Left
	}
	return next
}

func findTheDifference(s string, t string) byte {
	var m map[byte]int = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		m[s[i]-'a']--
	}
	for k, v := range m {
		if v != 0 {
			return k + 'a'
		}
	}
	return 0
}

var tmp []string
var f [][]bool
var res [][]string

func Partition(s string) [][]string {
	totalLen := len(s)
	f = make([][]bool, totalLen)
	res = make([][]string, 0)
	tmp = make([]string, 0)
	for i := 0; i < totalLen; i++ {
		f[i] = make([]bool, totalLen)
		f[i][i] = true

	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < totalLen; j++ {
			if s[i] == s[j] {

				if j > i+1 {
					f[i][j] = f[i+1][j-1]
				} else {
					f[i][j] = true
				}

			}
		}
	}
	dfs(s, 0)
	for _, v := range res {
		fmt.Println(v)
	}
	return res
}

func dfs(s string, i int) {
	if i == len(s) {
		newTmp := append([]string(nil), tmp...)
		res = append(res, newTmp)
		return
	}
	for j := i; j < len(s); j++ {
		if f[i][j] {
			tmp = append(tmp, s[i:j+1])
			dfs(s, j+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func maxScore(nums []int) int {

	nums2 := []int{1, 2, 3}

	nums2 = append(nums, 1)
	nums2 = append(nums, 2)

	fmt.Println(nums2)

	total := 0
	sort.Ints(nums)
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		if sum > 0 {
			total += 1
		} else {
			break
		}
	}

	return total
}

func minimumCost(cost []int) int {
	sort.Ints(cost)
	totalCost := 0
	i := len(cost) - 1
	for ; i >= 1; i = i - 3 {
		totalCost += (cost[i] + cost[i-1])
	}
	if i >= 0 {
		totalCost += cost[i]
	}
	return totalCost
}

func MakeFancyString(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); {
		c := s[i]
		j := i + 1
		for ; j < len(s); j++ {
			if c != s[j] {
				break
			}
		}
		end := min(j, i+2)
		if end > len(s) {
			end = len(s)
		}
		res.WriteString(s[i:end])
		i = j
	}
	return res.String()
}

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
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
