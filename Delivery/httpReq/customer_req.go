package httpReq

type CustomerReq struct {
	Name     string `json:"name"`
	Password string `json:"password" binding:"required"`
}
