package manager

import "payment/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
}

type useCaseManager struct {
	repoMag RepoManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoMag.CustomerRepo())
}

func NewUseCaseManager(repoMag RepoManager) UseCaseManager {
	return &useCaseManager{
		repoMag: repoMag,
	}
}
