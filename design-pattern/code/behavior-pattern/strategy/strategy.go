package strategy

import (
	"fmt"
)

// 定义一系列算法，让这些算法在运行时可以互换，使得算法分离，符合开闭原则

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name   string
	CardId string
	Money  int
}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Cash struct{}

type Bank struct{}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardId)
}
