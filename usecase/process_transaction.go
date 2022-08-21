package usecase

import (
	"github.com/Paulo-Lopes-Estevao/vitebank/domain"
	"github.com/Paulo-Lopes-Estevao/vitebank/dto"
	"time"
)

type UseCaseTransaction struct {
	transactionRepository domain.ITransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.ITransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{transactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(transactionDto)
	ccBalanceAndLimit, err := u.transactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance
	t := u.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)
	err = u.transactionRepository.SaveTransaction(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	transactionDto.ID = t.ID
	transactionDto.CreatedAt = t.CreatedAt
	return *t, nil
}

func (u UseCaseTransaction) newTransaction(transaction dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transaction.Amount
	t.Store = transaction.Store
	t.Description = transaction.Description
	t.CreatedAt = time.Now()
	return t
}

func (u UseCaseTransaction) hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}
