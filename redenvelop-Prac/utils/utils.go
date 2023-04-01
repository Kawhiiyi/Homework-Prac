package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"redenvelop-Prac/consts"
	"strings"
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

func StringToReader(s string) io.Reader {
	return strings.NewReader(s)
}

func Struct2Json(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
