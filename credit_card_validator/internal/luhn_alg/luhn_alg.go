package luhn_alg

type CreditCard struct {
	number []int
}

func NewCard(number []int) *CreditCard {
	return &CreditCard{
		number: number,
	}
}

func (c *CreditCard) Check() bool {
	total := 0
	checkSum := c.number[len(c.number)-1]

	for i := len(c.number) - 2; i >= 0; i-- {
		if i%2 == 0 {
			c.number[i] *= 2
			if c.number[i] > 9 {
				c.number[i] -= 9
			}
		}
		total += c.number[i]
	}

	return (total+checkSum)%10 == 0
}
