package main

import (
	"fmt"
	"sort"
)

func coinChange(money int, coinvalue []int) []int {
	sort.SliceStable(coinvalue, func(i, j int) bool {
		return coinvalue[i] > coinvalue[j]
	})
	result := []int{}
	for _, coin := range coinvalue {
		if money >= coin {
			for money >= coin {
				result = append(result, coin)
				money = money - coin
			}
		} else {
			continue
		}
	}

	return result
}

/*
koin: [10, 25, 5, 1]
uang: 42
42 = 1,1,1,1,1,1... (42x) --> 42
42 = 5,5,5,5,5,5,5,5,1,1 (10) -->42
42 = 10,10,10,10,1,1 (6) --> 42
42 = 10,10,10,5,5,1,1 (7) --> 42
42 = 25,5,5,5,1,1 (6) --> 42
42 = 25,10,5,1,1 (5) --> 42
*/

func main() {
	coins := []int{10, 25, 5, 1}
	fmt.Println(coinChange(13, coins))
}
