package usecase

import (
	"payment/Delivery/httpResp"
	"payment/entity"
	"payment/repository"
)

type TransferUseCase interface {
	TransferPayment(customerId, merchantId, cost int) (*entity.Transfer, error)
	ShowTransferDetail(transferId int) (*httpResp.TransferDetailResp, error)
}

type transferUseCase struct {
	repo repository.TransferRepo
}

func (t *transferUseCase) TransferPayment(customerId, merchantId, cost int) (*entity.Transfer, error) {
	return t.repo.Payment(customerId, merchantId, cost)
}

func (t *transferUseCase) ShowTransferDetail(transferId int) (*httpResp.TransferDetailResp, error) {
	return t.repo.GetDetail(transferId)
}

func NewTransferUseCase(repo repository.TransferRepo) TransferUseCase {
	return &transferUseCase{
		repo,
	}
}
