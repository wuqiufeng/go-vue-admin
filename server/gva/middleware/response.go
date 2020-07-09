package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResCode int

const (
	ResCodeOk                 = 1000
	ResCodeValidationError    = 2000
	ResCodeRequestParamsError = 2100
	ResCodeInnerServerError   = 5000
	ResCodeBizError           = 6000
)

type Response struct {
	Code    ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode, err error) {
	//resp := &Res{Code: code, Message: err.Error(), Data: make(map[string]interface{}, 0)}
	resp := &Response{Code: code, Message: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: ResCodeOk, Message: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
