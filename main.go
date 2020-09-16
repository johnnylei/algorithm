package main

import (
	"fmt"
	"github.com/johnnylei/algorithm/leet"
	_ "net/http/pprof"
)

type X int

func (x X) P()  {
	fmt.Println(x)
}

func main() {
	//x := X(10)
	//X.P(x)
	//graph := datastruct.InitDefault()
	//graph.String()
	//fmt.Println("BFS")
	//graph.BFS()
	//path := graph.MinPath("B", "E")
	//fmt.Println(path)
	//fmt.Println(dynamic_programming.Rob([]int{1, 3, 5, 7}))
	//// [[2,3],[4,5],[6,7],[8,9],[1,10]]
	////leet.Merge([][]int{{2,3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})
	////leet.LongestPalindrome("bsnetpqmwhqjunkooftuosgkmkxpmvuehtlpwpktltwlvpdaycnhewdbdrhluyjldecezujgxixehsmjjuyerpllrvzqskizkulqzowzfvqcdsllvgupndbaiuzihcxklvxbodpnrymwobhlvllybdlfabfvnizjpriapuzszdhohfgezayokrivbgbgingspoxsridokhwekawchjdcpylvefubulvxneuizglrjktfcdirwnpqztdpsicslzaeiaibrepifxpxfkczwoumkkuaqkbjhxvasxflmrctponwwenvmtdaqaavubyrzbqjbjxpejmucwunanxwpomqyondyjuzxmzpevxqmbkrwcpdiiph")
	////leet.ThreeSum([]int{-2,0,0,2,2})
	////fmt.Println(leet.WordBreak("ccaccc", []string{"cc","ac"}))
	////leet.BuildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	////fmt.Println(leet.GenerateParenthesis(3))
	////fmt.Println(leet.Search([]int{7,8,1,2,3,4,5,6}, 2))
	////leet.Permute([]int{5,4,6,2})
	//fmt.Println(leet.CombinationSum4([]int{1, 2, 3}, 4))
	////leet.GroupAnagrams([]string{"aaa", "bb", "cac", "acc", "cca"})
	////leet.SwapPairs(leet.InitList([]int{1, 2, 3, 4, 5}))
	////leet.WaysToChange(11)
	////fmt.Println(leet.CheckSubarraySum([]int{1, 2, 3}, 6))
	////leet.Reverse(leet.InitList([]int{1, 2, 3, 4, 5, 6}))
	////leet.PermuteUnique([]int{1, 1, 2})
	////leet.Largest1BorderedSquare([][]int{{0, 1}, {1, 1}})
	//leet.Largest1BorderedSquare([][]int{{1,1,0},{1,0,1},{1,1,1},{1,1,1},{1,1,1},{1,1,0},{1,1,1},{1,1,0}})

	//leet.CombinationSum([]int{2, 3, 5}, 8)
	//leet.Exist([][]byte{
	//	{'a', 'b'},
	//	{'c', 'd'},
	//}, "abcd")
	leet.Rob1([]int{2,1,1, 2})
}
