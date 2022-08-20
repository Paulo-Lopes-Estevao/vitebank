package repository

import (
	"github.com/Paulo-Lopes-Estevao/vitebank/domain"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	db *gorm.DB
}

func NewTransactionRepositoryDb(db *gorm.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, card domain.CreditCard) error {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	if transaction.Status == "" {
		t.updateBalance(card)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TransactionRepositoryDb) updateBalance(creditCard domain.CreditCard) error {
	err := t.db.Save(&creditCard).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) CreateCreditCard(creditCard domain.CreditCard) error {
	err := t.db.Create(&creditCard).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard
	err := t.db.Find(&c, "number=?", creditCard.Number).Error
	if err != nil {
		return c, err
	}
	return c, nil
}
