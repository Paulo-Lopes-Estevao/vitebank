package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type ITransactionRepository interface {
	SaveTransaction(transaction Transaction, card CreditCard) error
	GetCreditCard(card CreditCard) (CreditCard, error)
	CreateCreditCard(card CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t Transaction) ProcessAndValidate(card CreditCard) {
	if t.Amount+card.Balance > card.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		card.Balance = card.Balance + t.Amount
	}
}
