package manager

import "payment/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
	TransferRepo() repository.TransferRepo
	MerchantRepo() repository.MerchantRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.SqlDb())
}

func (r *repoManager) TransferRepo() repository.TransferRepo {
	return repository.NewTransferRepo(r.infra.SqlDb())
}

func (r *repoManager) MerchantRepo() repository.MerchantRepo {
	return repository.NewMerchantRepo(r.infra.SqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
