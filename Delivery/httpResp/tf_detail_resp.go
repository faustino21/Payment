package httpResp

import "time"

type TransferDetailResp struct {
	Id        int        `db:"transfer_id" json:"id,omitempty"`
	Customer  string     `db:"name" json:"customer,omitempty"`
	Merchant  string     `db:"merchant_name" json:"merchant,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}
