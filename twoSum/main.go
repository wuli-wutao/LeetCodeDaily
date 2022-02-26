package main

import (
	"fmt"
	"sort"
)

type Node struct {
	index int
	value int
}

func TwoSum(nums []int, target int) []int {
	ans := make([]int, 0)
	numsWithIndex := make([]Node, 0)

	for id, k := range nums {
		n := Node{
			index: id,
			value: k,
		}
		numsWithIndex = append(numsWithIndex, n)
	}

	sort.SliceStable(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i].value < numsWithIndex[j].value
	})

	i := 0
	j := len(nums) - 1

	for i < j {
		if numsWithIndex[i].value+numsWithIndex[j].value == target {
			ans = append(ans, numsWithIndex[i].index)
			ans = append(ans, numsWithIndex[j].index)
			break
		} else if numsWithIndex[i].value+numsWithIndex[j].value > target {
			j--
		} else {
			i++
		}
	}
	return ans
}

func main() {
	ans := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(ans)
}
