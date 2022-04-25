package manager

import "payment/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
	TransferUseCase() usecase.TransferUseCase
}

type useCaseManager struct {
	repoMag RepoManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoMag.CustomerRepo())
}

func (u *useCaseManager) TransferUseCase() usecase.TransferUseCase {
	return usecase.NewTransferUseCase(u.repoMag.TransferRepo())
}

func NewUseCaseManager(repoMag RepoManager) UseCaseManager {
	return &useCaseManager{
		repoMag: repoMag,
	}
}
