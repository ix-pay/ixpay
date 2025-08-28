package utils

import "github.com/gin-gonic/gin"

// type RespData[T any] struct {
//     Code int    `json:"code"`
//     Msg  string `json:"msg"`
//     Data T      `json:"data,omitempty"`
// }

type RespData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data" swaggertype:"object"`
}

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, &RespData{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, &RespData{
		Code: -1,
		Msg:  msg,
	})
}

func AbortError(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, &RespData{
		Code: -1,
		Msg:  msg,
	})
}
