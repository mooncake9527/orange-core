package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mooncake9527/orange-core/common/consts"
)

func GetReqId(c *gin.Context) string {
	reqId := c.GetString(consts.ReqId)
	if reqId == "" {
		reqId = uuid.NewString()
		c.Set(consts.ReqId, reqId)
	}
	return reqId
}
