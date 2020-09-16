package main

// 背包问题
// 1. 01背包问题，选或者不选
// 2. 完全背包问题
// 3. 多重背包问题
// 4. 混合背包问题
// 5. 二维费用的背包问题
// 6. 分组背包问题
// 7. 背包问题求方案数
// 8. 背包问题的方案
// 9. 有依赖的背包问题

type Goods struct {
	Volume int
	Value int
}

func Max(x, y int) int  {
	if x > y {
		return x
	}

	return y
}

// 最佳值
func Package01(goods []*Goods, volume int) int {
	goodsLength := len(goods)
	dp := make([][]int, goodsLength)
	for i := 0; i < goodsLength; i++ {
		dp[i] = make([]int, volume + 1)
		dp[i][0] = 0
		for v := 1; v <= volume; v++ {
			if i == 0 {
				if v >= goods[i].Volume {
					dp[i][v] = goods[i].Value
					continue
				}

				dp[i][v] = 0
				continue
			}

			if v < goods[i].Volume {
				dp[i][v] = dp[i - 1][v]
				continue
			}

			dp[i][v] = Max(dp[i - 1][v], dp[i - 1][v - goods[i].Volume] + goods[i].Value)
		}
	}
	return dp[goodsLength - 1][volume]
}

// 装了什么物品
// https://www.bilibili.com/video/BV1K4411X766?from=search&seid=17033201466856227946
// 前n个物品和前n-1个物品的最佳组合如果一致，说明，第n个物品没有被选择，反之，则代表第n个物品被选择了
func Package02(goods []*Goods, volume int) []*Goods {
	goodsLength := len(goods)
	dp := make([][]int, goodsLength)
	for i := 0; i < goodsLength; i++ {
		dp[i] = make([]int, volume + 1)
		dp[i][0] = 0
		for v := 1; v <= volume; v++ {
			if i == 0 {
				if v >= goods[i].Volume {
					dp[i][v] = goods[i].Value
					continue
				}

				dp[i][v] = 0
				continue
			}

			if v < goods[i].Volume {
				dp[i][v] = dp[i - 1][v]
				continue
			}

			dp[i][v] = Max(dp[i - 1][v], dp[i - 1][v - goods[i].Volume] + goods[i].Value)
		}
	}

	ret := make([]*Goods, 0)
	i, v := goodsLength - 1, volume
	for i >= 0 && v > 0 {
		if dp[i][v] == dp[i - 1][v] {
			i--
			continue
		}

		ret = append(ret, goods[i])
		v -= goods[i].Volume
		i--
	}
	return ret
}

// 可以重复, 最优解; 完全背包问题
func Package03(goods []*Goods, volume int) int {
	goodsLength := len(goods)
	dp := make([][]int, goodsLength)
	for i := 0; i < goodsLength; i++ {
		dp[i] = make([]int, volume + 1)
		for v := 0; v <= volume; v++ {
			if v == 0 {
				dp[i][v] = 0
				continue
			}

			if i == 0 {
				times := v / goods[i].Volume
				dp[i][v] = times * goods[i].Value
				continue
			}

			if v < goods[i].Volume {
				dp[i][v] = dp[i - 1][v]
				continue
			}

			dp[i][v] = Max(dp[i - 1][v], dp[i][v - goods[i].Volume] + goods[i].Value)
		}
	}
	return dp[goodsLength - 1][volume]
}

func main() {
	ret := Package02([]*Goods{
		&Goods{
			Volume:2,
			Value:3,
		},
		&Goods{
			Volume:3,
			Value:4,
		},
		&Goods{
			Volume:4,
			Value:5,
		},
		&Goods{
			Volume:5,
			Value:6,
		},
	}, 8)
	println(ret)
}
