package service

import (
	"my-wallet-ntier-mongo/repository"
	"my-wallet-ntier-mongo/response"
	userResponse "my-wallet-ntier-mongo/response"
	"my-wallet-ntier-mongo/utils/mapper"

	transactionResponse "my-wallet-ntier-mongo/response"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (u *TransactionService) GetTransactionTypes(queryParams response.TransactionTypeQuery) (response []transactionResponse.TransactionTypeResponse, total int64, err error) {
	transactionTypes := []userResponse.TransactionTypeResponse{}
	resp, total, err := u.transactionRepo.GetTransactionTypes(queryParams)
	if err != nil {
		return nil, 0, err
	}

	for _, transactionType := range resp {
		transactionTypes = append(transactionTypes, mapper.TransactionTypeModelToResponse(transactionType))
	}
	return transactionTypes, total, err
}

func (u *TransactionService) GetTransactionsByUserId(userId string, queryParams response.TransactionQuery) (response []transactionResponse.TransactionResponse, total int64, err error) {
	resp, total, err := u.transactionRepo.GetTransactionsByUserId(userId, queryParams)
	return resp, total, err
}
