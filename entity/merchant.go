package entity

import "time"

type Merchant struct {
	Id        int        `db:"merchant_id" json:"id,omitempty"`
	Name      string     `db:"merchant_name" json:"name,omitempty"`
	Category  string     `db:"category" json:"category,omitempty"`
	Saldo     int        `db:"saldo" json:"saldo,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
}
