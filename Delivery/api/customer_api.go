package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/Delivery/commonResp"
	"payment/Delivery/httpReq"
	"payment/Delivery/middleware"
	"payment/usecase"
	"payment/util"
	"strconv"
)

type CustomerApi struct {
	commonResp.BaseApi
	customer usecase.CustomerUseCase
}

func (cs *CustomerApi) VerifyLoginPasscode() gin.HandlerFunc {
	funcName := "LoginApi.VerifyLogin"

	return func(c *gin.Context) {
		var customer httpReq.CustomerReq

		err := cs.ParseRequestBody(c, &customer)
		if err != nil {
			cs.ParsingError(c, err)
			return
		}

		data, err := cs.customer.Authentication(customer.Name, customer.Password)
		if err != nil {
			util.LogError(funcName, "auth", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		token, err := middleware.GenerateToken(data.Name, data.Addres)
		if err != nil {
			util.LogError(funcName, "generateToken", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		err = cs.customer.InsertToken(token, data.Id)
		if err != nil {
			util.LogError(funcName, ".updateToken", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp(http.StatusOK, commonResp.NewSuccessMessage(gin.H{
			"customer": data,
			"token":    token,
		}))
	}
}

func (cs *CustomerApi) VerifyLogout() gin.HandlerFunc {
	funcName := "LoginApi.VerifyLogout"

	return func(c *gin.Context) {
		var customer httpReq.CustomerReq
		id, err := strconv.Atoi(c.Param("customerId"))
		if err != nil {
			util.LogError(funcName, ".parsingParam", err)
			return
		}
		err = cs.ParseRequestBody(c, &customer)
		if err != nil {
			cs.ParsingError(c, err)
			return
		}

		err = cs.customer.DeleteToken(id, customer.Password)
		if err != nil {
			util.LogError(funcName, ".deleteToken", err)
			commonResp.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, commonResp.NewFailedMessage(err.Error()))
			return
		}
		commonResp.NewAppHttpResponse(c).SuccessResp2(http.StatusOK, commonResp.NewSuccessMessage2())
	}
}

func LoginApiRoute(route *gin.RouterGroup, customer usecase.CustomerUseCase) *CustomerApi {
	customerApi := CustomerApi{
		customer: customer,
	}

	route.POST("login", customerApi.VerifyLoginPasscode())
	route.POST("/:customerId/logout", customerApi.VerifyLogout())
	return &customerApi
}
