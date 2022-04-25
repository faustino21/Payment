package manager

import "payment/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.SqlDb())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
