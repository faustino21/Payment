package usecase

import (
	"payment/entity"
	"payment/repository"
)

type TransferUseCase interface {
	TransferPayment(customerId, merchantId, cost int) (*entity.Transfer, error)
}

type transferUseCase struct {
	repo repository.TransferRepo
}

func (t *transferUseCase) TransferPayment(customerId, merchantId, cost int) (*entity.Transfer, error) {
	return t.repo.Payment(customerId, merchantId, cost)
}

func NewTransferUseCase(repo repository.TransferRepo) TransferUseCase {
	return &transferUseCase{
		repo,
	}
}
