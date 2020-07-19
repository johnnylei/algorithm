package main

import "fmt"

// document: https://www.bilibili.com/video/BV18x411V7fm

func fibonacciSequence(n int) int  {
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
func opt(arr []int) int {
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
func recSubset(arr []uint, current uint, target uint) bool {
	if current == 0 {
		return arr[current] == target
	}

	if target == arr[current] {
		return true
	}

	if target > arr[current] && recSubset(arr, current - 1, target - arr[current]) {
		return true
	}

	return recSubset(arr, current - 1, target)
}

func main() {
	fmt.Println(fibonacciSequence(5))
	fmt.Println(opt([]int{10, 8, 20, 19, 1, 8}))
	uarr := []uint{1, 3, 8, 9, 2, 0}
	fmt.Println(recSubset(uarr, uint(len(uarr) - 1), 17))
}