package commonResp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/util"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	err := c.ShouldBindJSON(body)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParsingError(c *gin.Context, err error) {
	util.LogError("Parsing", "", err)
	NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, NewFailedMessage(err.Error()))
}
