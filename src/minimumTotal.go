package main

import "fmt"

func main() {
	var tra = [][]int{{2},{3,4},{6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(tra))
}

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for k := range dp {
		dp[k] = make([]int, len(triangle[k]))
	}
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == len(triangle) - 1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j],dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
