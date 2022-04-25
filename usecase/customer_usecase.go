package usecase

import (
	"payment/entity"
	"payment/repository"
)

type CustomerUseCase interface {
	Authentication(username, password string) (*entity.Customer, error)
	InsertToken(token string, id int) error
	DeleteToken(id int, password string) error
}

type customerUseCaseImpl struct {
	repo repository.CustomerRepo
}

func (c *customerUseCaseImpl) Authentication(username, password string) (*entity.Customer, error) {
	return c.repo.Login(username, password)
}

func (c *customerUseCaseImpl) InsertToken(token string, id int) error {
	return c.repo.UpdateToken(token, id, "")
}

func (c *customerUseCaseImpl) DeleteToken(id int, password string) error {
	return c.repo.UpdateToken("", id, password)
}

func NewCustomerUseCase(repo repository.CustomerRepo) CustomerUseCase {
	return &customerUseCaseImpl{
		repo: repo,
	}
}
