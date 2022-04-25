package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"payment/Delivery/httpResp"
	"payment/entity"
	"payment/util"
)

type TransferRepo interface {
	Payment(customerId, merchantId, cost int) (*entity.Transfer, error)
	get(tx *sqlx.Tx, idTransfer int) (*entity.Transfer, error)
	Insert(tx *sqlx.Tx, customerId, merchantId, cost int) (int, error)
	GetDetail(transferId int) (*httpResp.TransferDetailResp, error)
}

type transferRepoImpl struct {
	db *sqlx.DB
}

func (t *transferRepoImpl) Payment(customerId, merchantId, cost int) (*entity.Transfer, error) {
	funcName := "TransferRepo.Insert"
	var transfer *entity.Transfer
	tx, err := t.db.Beginx()
	mrc := NewMerchantRepo(t.db)
	cst := NewCustomerRepo(t.db)

	defer func(tx *sqlx.Tx) {
		if err != nil {
			util.LogError(funcName, "", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}(tx)

	err = cst.UpdateSaldo(tx, customerId, cost)
	if err != nil {
		return nil, err
	}

	err = mrc.UpdateSaldo(tx, merchantId, cost)
	if err != nil {
		return nil, err
	}

	row, err := t.Insert(tx, customerId, merchantId, cost)
	transfer, err = t.get(tx, row)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}

func (t *transferRepoImpl) get(tx *sqlx.Tx, idTransfer int) (*entity.Transfer, error) {
	funcName := "TransferRepo.Get"
	var transfer entity.Transfer

	err := tx.Get(&transfer, "SELECT * FROM transfer WHERE transfer_id = ?", idTransfer)
	if err != nil {
		util.Log.Error().Msgf(funcName+".lasInsert : %v", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &transfer, nil
}

func (t *transferRepoImpl) Insert(tx *sqlx.Tx, customerId, merchantId, cost int) (int, error) {
	funcName := "TransferRepo.Insert"

	row := tx.MustExec("INSERT INTO transfer (id_customer, id_merchant, cost) VALUES (?,?,?)", customerId, merchantId, cost)
	if rowAffected, err := row.RowsAffected(); rowAffected == 0 || err != nil {
		util.LogError(funcName, ".rowsAffected", err)
		return 0, fmt.Errorf(err.Error())
	}
	id, _ := row.LastInsertId()
	intId := int(id)
	return intId, nil
}

func (t *transferRepoImpl) GetDetail(transferId int) (*httpResp.TransferDetailResp, error) {
	funcName := "TransferRepo.GetDetail"
	var transferDetail httpResp.TransferDetailResp

	err := t.db.Get(&transferDetail, "SELECT t.transfer_id, c.name, m.merchant_name, t.updated_at FROM transfer t INNER JOIN customers c ON t.id_customer = c.customer_id"+
		" INNER JOIN merchant m ON t.id_merchant = m.merchant_id WHERE t.transfer_id = ?", transferId)
	if err != nil {
		util.LogError(funcName, "", err)
		return nil, err
	}
	return &transferDetail, nil
}

func NewTransferRepo(db *sqlx.DB) TransferRepo {
	return &transferRepoImpl{
		db,
	}
}
