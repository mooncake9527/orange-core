package apis

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/npx/common/response"
)

type Api struct{}

type IError interface {
	Code() int
	Error() string
}

// Error 通常错误数据处理
func (e Api) Error(c *gin.Context, err error) {
	slog.Error("Error", "error", err)
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	if err, ok := err.(IError); ok {
		response.Error(c, err.Code(), err.Error())
		return
	}
	response.Error(c, 500, msg)
}

// ErrorC 通常错误数据处理
func (e Api) ErrorC(c *gin.Context, code int, msg string) {
	response.Error(c, code, msg)
}

// OK 通常成功数据处理
func (e Api) OK(c *gin.Context, data interface{}) {
	response.OK(c, data, "OK")
}

// PageOK 分页数据处理
func (e Api) PageOK(c *gin.Context, result interface{}, count int64, pageIndex int, pageSize int) {
	response.PageOK(c, result, count, pageIndex, pageSize, "OK")
}

// Custom 兼容函数
func (e Api) Custom(c *gin.Context, data gin.H) {
	response.Custum(c, data)
}
