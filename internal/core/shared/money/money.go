package money

type Money struct {
	amount int64
}

func NewMoney(a int64) Money {
	return Money{amount: a}
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Add(a Money) Money {
	return NewMoney(m.amount + a.amount)
}

func (m Money) Sub(a Money) Money {
	return NewMoney(m.amount - a.amount)
}

func (m Money) IsZero() bool {
	return m.amount == 0
}

func (m Money) IsNegative() bool {
	return m.amount < 0
}

func (m Money) IsGreaterThan(v Money) bool {
	return m.amount > v.amount
}
