package main

import (
	"errors"
)

var errNegativeBalance = errors.New("ошибка: баланс не может быть отрицательным")
var errNegativeQuantity = errors.New("ошибка: количество не может быть отрицательным")

type Account struct {
	owner   string
	balance float64
}

// конструктор Account
func NewAccount(owner string) *Account {
	return &Account{
		owner: owner,
	}
}

// сеттер balance
func (a *Account) SetBalance(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	a.balance = quantity
	return nil
}

// геттер balance
func (a *Account) GetBalance() float64 {
	return a.balance
}

// метод зачисления средств на счет
func (a *Account) Deposit(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	a.balance += quantity
	return nil
}

// метод снятия средств со счета
func (a *Account) Withdraw(quantity float64) error {
	if quantity < 0 {
		return errNegativeQuantity
	}

	if a.balance-quantity < 0 {
		return errNegativeBalance
	}

	a.balance -= quantity
	return nil
}
