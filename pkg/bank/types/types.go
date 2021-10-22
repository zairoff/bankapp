package types

type Paymentsource struct {
	Type    string
	Number  string
	Balance Money
}

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
	Number   string
	Type     string
}

type Payment struct {
	ID     int
	Amount Money
}
