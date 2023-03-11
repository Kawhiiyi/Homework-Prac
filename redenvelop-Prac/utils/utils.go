package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redenvelop-Prac/consts"
)

func RetJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.Success.Code,
		"msg":  consts.Success.Msg,
	})
}

func RetJsonWithData(c *gin.Context, data string) {
	c.JSON(http.StatusOK, gin.H{
		"code": consts.Success.Code,
		"msg":  consts.Success.Msg,
		"data": data,
	})
}

func RetErrJson(c *gin.Context, rErr consts.RError) {
	c.JSON(http.StatusOK, gin.H{
		"code": rErr.Code,
		"msg":  rErr.Msg,
	})
}
