package entity

import (
	"time"
)

type Customer struct {
	Id        int        `db:"customer_id" json:"id,omitempty"`
	Name      string     `db:"name" json:"name,omitempty"`
	Password  string     `db:"password" json:"password,omitempty"`
	Addres    string     `db:"address" json:"addres,omitempty"`
	Saldo     int        `db:"saldo" json:"saldo,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt,omitempty"`
}
