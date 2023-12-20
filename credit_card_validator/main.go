package main

import (
	"credit_card_validator/internal/luhn_alg"
	"fmt"
)

func main() {
	c := luhn_alg.NewCard([]int{4, 5, 3, 9, 6, 7, 7, 9, 0, 8, 0, 1, 6, 8, 0, 8})

	fmt.Println(c.Check())
}
