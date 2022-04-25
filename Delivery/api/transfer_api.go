package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/Delivery/commonResp"
	"payment/Delivery/httpReq"
	"payment/usecase"
	"payment/util"
	"strconv"
)

type TransferApi struct {
	commonResp.BaseApi
	transfer usecase.TransferUseCase
}

func (t *TransferApi) Payment() gin.HandlerFunc {
	funcName := "TransferApi.Payment"

	return func(c *gin.Context) {
		var payment httpReq.TransferReq

		customerId, err := strconv.Atoi(c.Param("customerId"))
		if err != nil {
			util.LogError(funcName, ".customerId", err)
			return
		}
		merchantId, err := strconv.Atoi(c.Param("merchantId"))
		if err != nil {
			util.LogError(funcName, ".merchantId", err)
			return
		}

		err = t.ParseRequestBody(c, &payment)
		if err != nil {
			t.ParsingError(c, err)
			return
		}

		data, err := t.transfer.TransferPayment(customerId, merchantId, payment.Cost)
		if err != nil {
			util.LogError(funcName, "", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(data))
	}
}

func PaymentApi(route *gin.RouterGroup, transfer usecase.TransferUseCase) *TransferApi {
	transferApi := TransferApi{
		transfer: transfer,
	}

	route.POST("/:customerId/:merchantId", transferApi.Payment())
	return &transferApi
}
