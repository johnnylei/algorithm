package leet

import (
	"github.com/johnnylei/algorithm/datastruct"
	"sort"
	"strconv"
	"strings"
)

// 56, 区间合并
func Merge(intervals [][]int) [][]int {
	yLen := len(intervals)
	if yLen == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0)
	for i := 0; i < yLen; {
		item := intervals[i]
		j := i + 1
		for ; j < yLen; j++ {
			if item[1] < intervals[j][0] {
				break
			}

			if item[1] < intervals[j][1] {
				item[1] = intervals[j][1]
			}
		}

		i = j
		ret = append(ret, item)
	}

	return ret
}

// 5, 最长回文，动态规划解法
func LongestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	max := s[0 : 1]
	dp := make([][]bool, length)
	for i := length - 2; i >= 0; i-- {
		dp[i] = make([]bool, length)
		dp[i][i] = true
		for j := i + 1; j < length; j++ {
			if j == i + 1 {
				dp[i][j] = s[i] == s[j]
				if dp[i][j] && len(max) < 2 {
					max = s[i : j + 1]
				}
				continue
			}

			dp[i][j] = dp[i + 1][j - 1] && s[i] == s[j]
			if dp[i][j] && len(max) < j + 1 - i {
				max = s[i : j + 1]
			}
		}
	}

	return max
}

// 647. 回文子串
func countSubstrings(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	dp := make([][]bool, length)
	times := length
	for i := length - 2; i >= 0; i-- {
		dp[i] = make([]bool, length)
		dp[i][i] = true
		for j := i + 1; j < length; j++ {
			if j == i + 1 {
				dp[i][j] = s[i] == s[j]
				if dp[i][j] {
					times++
				}
				continue
			}

			dp[i][j] = dp[i + 1][j - 1] && s[i] == s[j]
			if dp[i][j] {
				times++
			}
		}
	}
	return times
}

// 15, 三数之和
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	if nums[length - 1] < 0 {
		return nil
	}

	ret := make([][]int, 0)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i - 1] {
			break
		}

		for j := length - 1; j > i; j-- {
			if nums[j] < 0 {
				break
			}

			if j < length - 1 && nums[j] == nums[j + 1] {
				continue
			}

			for z := i + 1; z < j; z++ {
				if nums[i] + nums[j] + nums[z] == 0 {
					ret = append(ret, []int{nums[i], nums[z], nums[j]})
					break
				}
			}
		}
	}

	return ret
}

// 139； s 是否能够通过 wordDict组合生成
func WordBreak(s string, wordDict []string) bool {
	picked := make([]string, 0)
	for i := 0; i < len(wordDict); i++ {
		if strings.Contains(s, wordDict[i]) {
			picked = append(picked, wordDict[i])
		}
	}

	find := func(s string, wordDict []string, dp []bool) bool {
		sLength := len(s)
		for _, item := range wordDict {
			itemLen := len(item)
			if itemLen > sLength {
				continue
			}

			if !strings.Contains(s, item) {
				continue
			}

			if item == s {
				return true
			}

			if !dp[sLength - itemLen - 1] {
				continue
			}

			if item == s[sLength - itemLen:] {
				return true
			}
		}

		return false
	}

	length := len(s)
	dp := make([]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = find(s[0 : i + 1], picked, dp)
	}

	return dp[length - 1]
}

// 105
func BuildTree(preorder []int, inorder []int) *datastruct.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &datastruct.TreeNode{
		Val:preorder[0],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = BuildTree(preorder[1 : i + 1], inorder[:i])
	root.Right = BuildTree(preorder[i + 1:], inorder[i + 1:])
	return root
}

// 106， 通过先序中序遍历，生成树
func buildTree(inorder []int, postorder []int) *datastruct.TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return &datastruct.TreeNode{
			Val:inorder[0],
		}
	}

	root := &datastruct.TreeNode{
		Val: postorder[length - 1],
	}
	i := 0
	for ; i < length; i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i + 1:], postorder[i: length - 1])
	return root
}

// 22 括号生成
func GenerateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	if n == 1 {
		return []string{
			"()",
		}
	}

	sub := GenerateParenthesis(n - 1)
	ret := make([]string, 0)
	resultMap := make(map[string]bool)
	for _, subItem := range sub {
		for i := 0; i < len(subItem); i++ {
			item := subItem[0:i] + "()" + subItem[i:]
			if resultMap[item] {
				continue
			}

			ret = append(ret, item)
			resultMap[item] = true
		}
	}
	return ret
}

// 33; 旋转数组，二分查找
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	type SearchFunc func([]int, int, int, int) int
	var searchFunc SearchFunc
	searchFunc = func(nums []int, start int, end int, target int) int {
		length := end - start
		if length <= 0 {
			return -1
		}

		middle := length / 2 + start
		if target == nums[middle] {
			return middle
		}

		// 有序的情况
		if nums[start] < nums[end - 1] {
			if target > nums[middle] {
				return searchFunc(nums, middle + 1, end, target)
			}

			return searchFunc(nums, start, middle, target)
		}

		// 前段有序, 后段无序
		if nums[start] < nums[middle] {
			if target <= nums[middle] && target >= nums[start] {
				return searchFunc(nums, start, middle, target)
			}

			return searchFunc(nums, middle + 1, end, target)
		}

		// 后段有序， 前段无序
		if target >= nums[middle] && target <= nums[end - 1] {
			return searchFunc(nums, middle + 1, end, target)
		}

		return searchFunc(nums, start, middle, target)
	}

	return searchFunc(nums, 0, len(nums), target)
}

// 46， 全排列
func Permute1(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	result := make([][][]int, 0)
	for i := 0; i < length; i++ {
		if i == 0 {
			result = append(result, [][]int{
				{
					nums[0],
				},
			})
			continue
		}

		iResult := make([][]int, 0)
		for _, subResult := range result[i - 1] {
			for j := 0; j <= len(subResult); j++ {
				item := make([]int, 0)
				item = append(item, subResult[:j]...)
				item = append(item, nums[i])
				item = append(item, subResult[j:]...)
				iResult = append(iResult, item)
			}
		}
		result = append(result, iResult)
	}

	return result[length - 1]
}

// 46， 全排列; 回溯算法，官方解法
func Permute(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	used := make([]bool, length)
	path := make([]int, 0)
	result := make([][]int, 0)
	type DFS func(nums []int, depth int, used *[]bool, path *[]int, result *[][]int)
	var dfs DFS
	dfs = func(nums []int, depth int, used *[]bool, path *[]int, result *[][]int) {
		if depth == length {
			item := make([]int, length)
			copy(item, *path)
			*result = append(*result, item)
			return
		}

		for i := 0; i < length; i++ {
			if (*used)[i] {
				continue
			}

			(*used)[i] = true
			depth++
			*path = append(*path, nums[i])
			dfs(nums, depth, used, path, result)
			*path = (*path)[:len(*path) - 1]
			depth--
			(*used)[i] = false
		}
	}
	dfs(nums, 0, &used, &path, &result)
	return result
}

// 47, 去重全排列
func PermuteUnique(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	sort.Ints(nums)
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, length)
	type DFS func(nums []int, depth int, path *[]int, used *[]bool, result *[][]int)
	var dfs DFS
	dfs = func(nums []int, depth int, path *[]int, used *[]bool, result *[][]int) {
		if depth == length {
			item := make([]int, length)
			copy(item, *path)
			*result = append(*result, item)
			return
		}

		for i := 0; i < length; i++ {
			if (*used)[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i - 1] && !(*used)[i - 1] {
				continue
			}

			depth++
			*path = append(*path, nums[i])
			(*used)[i] = true
			dfs(nums, depth, path, used, result)
			(*used)[i] = false
			*path = (*path)[:len(*path) - 1]
			depth--
		}
	}

	dfs(nums, 0, &path, &used, &result)
	return result
}

// 60, 第k个排列
func GetPermutation(n int, k int) string {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	ret := ""
	used := make([]bool, n)
	path := ""
	type DFS func(nums []int, depth int, used []bool, path string, k *int) bool
	var dfs DFS
	dfs = func(nums []int, depth int, used []bool, path string, k *int) bool {
		if depth == n {
			*k--
			if *k == 0 {
				ret = path
			}
			return *k == 0
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path += strconv.Itoa(nums[i])
			depth++
			if dfs(nums, depth, used, path, k) {
				return true
			}
			depth--
			path = path[:len(path) - 1]
			used[i] = false
		}

		return false
	}
	dfs(nums, 0, used, path, &k)
	return ret
}

// 39; 深度优先遍历; 解决candidates组合和等于target问题
func CombinationSum(candidates []int, target int) [][]int {
	length := len(candidates)
	if length == 0 {
		return nil
	}

	type DFS func(candidates []int, target int, start int, path *[]int, result *[][]int)
	var dfs DFS
	dfs = func(candidates []int, target int, start int, path *[]int, result *[][]int) {
		if target == 0 {
			item := make([]int, len(*path))
			copy(item, *path)
			*result = append(*result, item)
			return
		}

		for i := start; i < length; i++ {
			if target < candidates[i] {
				continue
			}

			*path = append(*path, candidates[i])
			dfs(candidates, target - candidates[i], i, path, result)
			*path = (*path)[:len(*path) - 1]
		}
	}

	result := make([][]int, 0)
	path := make([]int, 0)
	dfs(candidates, target, 0, &path, &result)
	return result
}

// 40 数组总和2，
func combinationSum2(candidates []int, target int) [][]int {
	length := len(candidates)
	if length == 0 {
		return nil
	}

	sort.Ints(candidates)
	type DFS func([]int, int, *[]int, *[][]int)
	var dfs DFS
	dfs = func(candidates []int, target int, path *[]int, result *[][]int) {
		if target == 0 {
			item := make([]int, len(*path))
			copy(item, *path)
			*result = append(*result, item)
			return
		}

		length := len(candidates)
		for i := 0; i < length; i++ {
			if target < candidates[i] {
				continue
			}

			if i > 0 && candidates[i] == candidates[i - 1] {
				continue
			}

			*path = append(*path, candidates[i])
			subCandidate := candidates[i+1:]
			dfs(subCandidate, target - candidates[i], path, result)
			*path = (*path)[:len(*path) - 1]
		}
	}

	path := make([]int, 0)
	result := make([][]int, 0)
	dfs(candidates, target, &path, &result)
	return result
}

// 377 组合总和
func CombinationSum4(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make([]int, target + 1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < length; j++ {
			if i >= nums[j] {
				dp[i] += dp[i - nums[j]]
			}
		}
	}

	return dp[target]
}

// 49；给定一个字符串，将其按照字符组成分组
func GroupAnagrams(strs []string) [][]string {
	length := len(strs)
	if length == 0 {
		return nil
	}

	type QuickSort func(arr *[]byte, start, end int)
	var sort QuickSort
	sort = func(arr *[]byte, start, end int) {
		if start >= end {
			return
		}

		left, right, pointer := start, end, (*arr)[start]
		for left < right {
			if (*arr)[left] > (*arr)[right] {
				(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
			}

			if pointer == (*arr)[left] {
				right--
			} else {
				left++
			}
		}

		if start < left - 1 {
			sort(arr, start, left - 1)
		}

		if end > left + 1 {
			sort(arr, left + 1, end)
		}
	}

	strMap := make(map[string][]string)
	for _, str := range strs {
		strBytes := []byte(str)
		sort(&strBytes, 0, len(strBytes) - 1)
		strMap[string(strBytes)] = append(strMap[string(strBytes)], str)
	}

	result := make([][]string, 0)
	for _, item := range strMap {
		result = append(result, item)
	}

	return result
}


// 48，n*n矩阵旋转
func rotate(matrix [][]int)  {
	length := len(matrix)
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length / 2; j++ {
			matrix[i][j], matrix[i][length - 1 - j] = matrix[i][length - 1 - j], matrix[i][j]
		}
	}
}

// 24, 链表中的节点两两交换
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current, next := head, head.Next
	if next.Next != nil {
		next.Next = SwapPairs(next.Next)
	}
	current.Next = next.Next
	next.Next = current
	return next
}

// 08.11. 硬币； 1000000007
func waysToChange(n int) int {
	dp := make([][]int, 4)
	coins := []int{
		1, 5, 10, 25,
	}

	for i := 0; i < 4; i++ {
		dp[i] = make([]int, n + 1)
		dp[i][0] = 1
		for j := 1; j <= n; j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}

			metric := coins[i]
			if j < metric {
				dp[i][j] = dp[i - 1][j]
				continue
			}

			dp[i][j] = dp[i - 1][j] + dp[i][j - metric]
		}
	}

	return dp[3][n] % 1000000007
}

// O-42. 连续子数组的最大和
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max := nums[0]
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}

		if dp[i - 1] > 0 {
			dp[i] = dp[i - 1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}


// 61； 往右边移动k个节点
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	current, length := head, 0
	for current.Next != nil {
		length++
		current = current.Next
	}

	current.Next = head
	k = k % length
	current = head
	for i := 0; i < k - 1; i++ {
		current = current.Next
	}
	head = current.Next
	current.Next = nil
	return head
}

// 11. 盛最多水的容器
func MaxArea(height []int) int {
	maxArea := 0
	h := 0
	length := len(height)
	i, j := 0, length - 1
	for i < j {
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}
		w := j + 1 - i
		area := h * w
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// 62; 不同路径
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}

			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}

	return dp[m - 1][n - 1]
}

// 63, 不同路径2
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	yLength := len(obstacleGrid)
	if yLength == 0 {
		return 0
	}
	
	xLength := len(obstacleGrid[0])
	if xLength == 0 {
		return 0
	}
	
	dp := make([][]int, yLength)
	for i := 0; i < yLength; i++ {
		dp[i] = make([]int, xLength)
		for j := 0; j < xLength; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			
			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			
			if i == 0 {
				dp[i][j] = dp[i][j - 1]
				continue
			}
			
			if j == 0 {
				dp[i][j] = dp[i - 1][j]
				continue
			}
			
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	return dp[yLength - 1][xLength - 1]
}

// 1139. 最大的以 1 为边界的正方形
func largest1BorderedSquare(grid [][]int) int {
	yLen := len(grid)
	if yLen == 0 {
		return 0
	}

	xLen := len(grid[0])
	if xLen == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}

		return x
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}

		return x
	}

	maxLength := 0
	dp := make([][][]int, yLen)
	for y := 0; y < yLen; y++ {
		dp[y] = make([][]int, xLen)
		for x := 0; x < xLen; x++ {
			dp[y][x] = make([]int, 2)
			if grid[y][x] != 1 {
				continue
			}

			if maxLength == 0 {
				maxLength = 1
			}

			if x == 0 && y == 0 {
				dp[y][x][0] = 1
				dp[y][x][1] = 1
				continue
			}

			if y == 0 {
				dp[y][x][0] += dp[y][x - 1][0] + 1
				dp[y][x][1] = 1
				continue
			}

			if x == 0 {
				dp[y][x][0] = 1
				dp[y][x][1] += dp[y - 1][x][1] + 1
				continue
			}

			dp[y][x][0] += dp[y][x - 1][0] + 1
			dp[y][x][1] += dp[y - 1][x][1] + 1
			length := min(dp[y][x][0], dp[y][x][1])
			for length > 0 {
				if dp[y][x + 1 - length][1] >= length && dp[y + 1 - length][x][0] >= length {
					break
				}
				length--
			}
			maxLength = max(maxLength, length)
		}
	}
	return maxLength * maxLength
}

// 79. 单词搜索
func Exist(board [][]byte, word string) bool {
	wordLen := len(word)
	yLength := len(board)
	xLength := len(board[0])

	wordMap := make(map[byte]int, wordLen)
	for i := 0 ; i < wordLen; i++ {
		wordMap[word[i]] += 1
	}
	
	for i := 0; i < yLength; i++ {
		for j := 0; j < xLength; j++ {
			if val, okay := wordMap[board[i][j]]; okay && val > 0 {
				wordMap[board[i][j]] -= 1
				wordLen--
				if wordLen == 0 {
					return true
				}
			}
		}
	}

	return false
}

// 198. 打家劫舍
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	dp := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[0] = nums[0]
			continue
		}

		if i == 1 {
			dp[i] = max(dp[0], nums[1])
			continue
		}

		dp[i] = max(dp[i - 1], nums[i] + dp[i - 2])
	}

	return dp[length - 1]
}

// 213. 打家劫舍 II
func Rob1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	// 不偷第一个房间
	dp := make([]int, length)
	dp[0] = 0
	for i := 1; i < length; i++ {
		if i == 1 {
			dp[i] = nums[1]
			continue
		}

		dp[i] = max(dp[i - 1], nums[i] + dp[i - 2])
	}

	maxRet := dp[length - 1]
	for i := 0; i < length - 1; i++ {
		if i == 0 {
			dp[0] = nums[0]
			continue
		}

		if i == 1 {
			dp[i] = max(dp[0], nums[1])
			continue
		}

		dp[i] = max(dp[i - 1], nums[i] + dp[i - 2])
	}

	return max(maxRet, dp[length - 2])
}
