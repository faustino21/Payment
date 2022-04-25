package entity

import "time"

type Transfer struct {
	Id         int        `db:"transfer_id" json:"id,omitempty"`
	CustomerId int        `db:"id_customer" json:"customer,omitempty"`
	MerchantId int        `db:"id_merchant" json:"merchantId,omitempty"`
	Cost       int        `db:"cost" json:"cost,omitempty"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	CreatedAt  *time.Time `db:"created_at" json:"createdAt,omitempty"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}
