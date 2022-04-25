package httpReq

type TransferReq struct {
	Cost int `json:"cost"`
}

type TransferDetailReq struct {
	TransferId int `json:"transferId"`
}
