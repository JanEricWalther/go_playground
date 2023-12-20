package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		sum      int
		expected int
	}{
		{
			[]int{1, 2, 3},
			4,
			4,
		},
		{
			[]int{2, 5, 3, 6},
			10,
			5,
		},
	}

	for _, test := range tests {
		actual := CoinChangeRecursive(test.coins, test.sum)
		if actual != test.expected {
			t.Errorf("Recursive (%v, %d): expected %d, got %d",
				test.coins, test.sum, test.expected, actual)
		}
		actual = CoinChangeDynamic(test.coins, test.sum)
		if actual != test.expected {
			t.Errorf("Dynamic (%v, %d): expected %d, got %d",
				test.coins, test.sum, test.expected, actual)
		}
	}
}

func BenchmarkRecursive(b *testing.B) {
	coins := []int{1, 3, 7, 13}
	sum := 2560
	for i := 0; i < b.N; i++ {
		CoinChangeRecursive(coins, sum)
	}
}

func BenchmarkDynamic(b *testing.B) {
	coins := []int{1, 3, 7, 13}
	sum := 2560
	for i := 0; i < b.N; i++ {
		CoinChangeDynamic(coins, sum)
	}
}
