package main

import "fmt"

type Money struct {
	Balance int
}

func (m *Money) Add(amount int) {
	m.Balance += amount
}

func (m *Money) Mul(amount int) {
	m.Balance -= amount
}

func (m *Money) String() string {
	return fmt.Sprintf("Current balance: %d\n", m.Balance)
}

func NewMoney(balance int) *Money {
	return &Money{Balance: balance}
}

func main() {
	m := NewMoney(100)
	fmt.Println(m.String())

	m.Add(23)
	fmt.Println(m.String())

	m.Mul(100)
	fmt.Println(m.String())

}
