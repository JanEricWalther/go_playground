package coinchange

func CoinChangeRecursive(coins []int, sum int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 {
		return 0
	}

	if len(coins) == 0 {
		return 0
	}

	return CoinChangeRecursive(coins[1:], sum) + CoinChangeRecursive(coins, sum-coins[0])
}

func CoinChangeDynamic(coins []int, sum int) int {
	len := len(coins)
	table := make([][]int, len+1)
	for i := range table {
		table[i] = make([]int, sum+1)
	}

	for i := 0; i <= len; i++ {
		table[i][0] = 1
	}

	for i := 1; i <= len; i++ {
		for j := 1; j <= sum; j++ {
			if coins[i-1] > j {
				table[i][j] = table[i-1][j]
			} else {
				table[i][j] = table[i-1][j] + table[i][j-coins[i-1]]
			}
		}
	}

	return table[len][sum]
}
