package transaction

import "github.com/adrianostankewicz/bank/internal/core/shared/money"

func CalculateBalance(transactions []*Transaction) money.Money {
	total := money.NewMoney(0)
	for _, t := range transactions {
		total = total.Add(t.Amount())
	}
	return total
}
