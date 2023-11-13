package designpattern

import "fmt"

// 目标接口：统一的支付接口
type Payment interface {
	ProcessPayment(amount float64) string
}

// 适配器：将不同的支付方式适配为统一的 Payment 接口
type CreditCardAdapter struct {
	CreditCard *CreditCard
}

func (a *CreditCardAdapter) ProcessPayment(amount float64) string {
	return a.CreditCard.Charge(amount)
}

type PayPalAdapter struct {
	PayPal *PayPal
}

func (a *PayPalAdapter) ProcessPayment(amount float64) string {
	return a.PayPal.ExecutePayment(amount)
}

// 被适配的类：具有不同接口的支付方式
type CreditCard struct{}

func (c *CreditCard) Charge(amount float64) string {
	return fmt.Sprintf("Charging %v using Credit Card", amount)
}

type PayPal struct{}

func (p *PayPal) ExecutePayment(amount float64) string {
	return fmt.Sprintf("Executing payment of %v using PayPal", amount)
}
