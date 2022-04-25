package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"payment/entity"
	"payment/util"
)

type CustomerRepo interface {
	Login(username, password string) (*entity.Customer, error)
	UpdateToken(token string, id int, password string) error
}

type customerRepoImpl struct {
	db *sqlx.DB
}

func (c *customerRepoImpl) Login(username, password string) (*entity.Customer, error) {
	funcName := "CustomerRepo.Login"

	var customer entity.Customer
	err := c.db.Get(&customer, "SELECT customer_id, name, address, updated_at, created_at FROM customers WHERE name = ? AND password = ?", username, password)
	if err != nil {
		util.LogError(funcName, "getQuery", err)
		return nil, fmt.Errorf(err.Error())
	}
	if customer.Id == 0 {
		util.LogError(funcName, "", fmt.Errorf("Not have customer id"))
		return nil, fmt.Errorf("unauthorized")
	}
	return &customer, nil
}

func (c *customerRepoImpl) UpdateToken(token string, id int, password string) error {
	funcName := "CustomerRepo.UpdateToken"

	tx := c.db.MustBegin()

	if password == "" {
		row := tx.MustExec("UPDATE customers SET token = ? WHERE customer_id = ?", token, id)
		if x, err := row.RowsAffected(); x == 0 {
			util.LogError(funcName, ".Exec Update", err)
		}
	} else {
		row := tx.MustExec("UPDATE customers SET token = null WHERE customer_id = ? AND password = ? ", id, password)
		if x, err := row.RowsAffected(); x == 0 {
			util.LogError(funcName, ".Exec Update", err)
			return fmt.Errorf("user not found")
		}
	}
	err := tx.Commit()
	if err != nil {
		util.LogError(funcName, ".commit", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepo {
	return &customerRepoImpl{
		db,
	}
}
