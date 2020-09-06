package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 成功响应
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "SUCCESS",
		"data": data,
	})
}

// 失败响应
func JsonFail(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": msg,
	})
}

// 失败响应并返回数据
func JsonFailData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": msg,
		"data": data,
	})
}

// 错误响应
func JsonError(ctx *gin.Context, code int, msg string)  {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
	})
}

// 错误响应并返回数据
func JsonErrorData(ctx *gin.Context, code int, msg string, data interface{})  {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}
