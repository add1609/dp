package stairs

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		if dp[i-2] < dp[i-1] {
			dp[i] = cost[i] + dp[i-2]
		} else {
			dp[i] = cost[i] + dp[i-1]
		}
	}
	if dp[len(dp)-2] < dp[len(dp)-1] {
		return dp[len(dp)-2]
	}
	return dp[len(dp)-1]
}
