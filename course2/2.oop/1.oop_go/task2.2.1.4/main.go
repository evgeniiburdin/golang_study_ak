package main

import (
	"errors"
	"fmt"
)

// PaymentMethod интерфейс определяет метод Pay для обработки платежа
type PaymentMethod interface {
	Pay(amount float64) error
}

// CreditCard структура представляет кредитную карту
type CreditCard struct {
	balance float64
}

// Pay метод для CreditCard
func (cc *CreditCard) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("недопустимая сумма платежа")
	}
	if cc.balance < amount {
		return errors.New("недостаточный баланс")
	}
	cc.balance -= amount
	fmt.Printf("Оплачено %.2f с помощью кредитной карты\n", amount)
	return nil
}

// Bitcoin структура представляет биткоин
type Bitcoin struct {
	balance float64
}

// Pay метод для Bitcoin
func (btc *Bitcoin) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("недопустимая сумма платежа")
	}
	if btc.balance < amount {
		return errors.New("недостаточный баланс")
	}
	btc.balance -= amount
	fmt.Printf("Оплачено %.2f с помощью биткоина\n", amount)
	return nil
}

// ProcessPayment функция для обработки платежа
func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProcessPayment(cc, 200.00) // Оплачено 200.00 с помощью кредитной карты
	ProcessPayment(btc, 1.00)  // Оплачено 1.00 с помощью биткоина
	ProcessPayment(cc, 0.00)   // Не удалось обработать платеж: недопустимая сумма платежа
	ProcessPayment(cc, 600.00) // Не удалось обработать платеж: недостаточный баланс
	ProcessPayment(btc, 3.00)  // Не удалось обработать платеж: недостаточный баланс
}
