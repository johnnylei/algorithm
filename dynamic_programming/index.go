package dynamic_programming

// document: https://www.bilibili.com/video/BV18x411V7fm

func FibonacciSequence(n int) int  {
	if n < 3 {
		return 1
	}

	fn, fn_1, fn_2 := 0, 1, 1
	for i := 2; i < n; i++ {
		fn = fn_1 + fn_2
		fn_1 = fn_2
		fn_2 = fn
	}
	return fn
}

// 相邻的不相加，求最大值
func Opt(arr []int) int {
	length := len(arr)
	if length == 0 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make(map[int]int, 0)
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[0] = arr[0]
			continue
		}

		if i == 1 {
			dp[1] = max(arr[0], arr[1])
			continue
		}

		dp[i] = max(dp[i - 1], dp[i - 2] + arr[i])
	}
	return dp[length - 1]
}

// 一组数组，是否存在目标值
func RecSubset(arr []uint, current uint, target uint) bool {
	if current == 0 {
		return arr[current] == target
	}

	if target == arr[current] {
		return true
	}

	if target > arr[current] && RecSubset(arr, current - 1, target - arr[current]) {
		return true
	}

	return RecSubset(arr, current - 1, target)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// leetcode: 198
func Rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make(map[int]int)
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[i] = nums[0]
			continue
		}

		if i == 1 {
			dp[i] = max(nums[0], nums[1])
			continue
		}

		dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	}

	return dp[length - 1]
}

