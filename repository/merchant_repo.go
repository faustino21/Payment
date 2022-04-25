package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"payment/entity"
	"payment/util"
)

type MerchantRepo interface {
	GetMerchant(tx *sqlx.Tx, id int) (*entity.Merchant, error)
	UpdateSaldo(tx *sqlx.Tx, id, upSaldo int) error
}

type merchantRepoImpl struct {
	db *sqlx.DB
}

func (m *merchantRepoImpl) GetMerchant(tx *sqlx.Tx, id int) (*entity.Merchant, error) {
	funcName := "MerchantRepo.GetMerchant"
	var merchant entity.Merchant

	err := tx.Get(&merchant, "SELECT merchant_name, saldo FROM merchant WHERE merchant_id = ?", id)
	if err != nil {
		util.LogError(funcName, "", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &merchant, nil
}

func (m *merchantRepoImpl) UpdateSaldo(tx *sqlx.Tx, id, upSaldo int) error {
	funcName := "MerchantRepo.UpdateSaldo"
	merchant, err := m.GetMerchant(tx, id)
	if err != nil {
		return err
	}
	cost := merchant.Saldo + upSaldo
	row, err := tx.Exec("UPDATE merchant SET saldo = ? WHERE merchant_id = ? ", cost, id)
	rowAffected, err := row.RowsAffected()
	if rowAffected == 0 && err != nil {
		util.LogError(funcName, ".rowsAffected", err)
		return err
	}
	return nil
}

func NewMerchantRepo(db *sqlx.DB) MerchantRepo {
	return &merchantRepoImpl{
		db,
	}
}
