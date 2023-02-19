package main //持久化存储
import (
	"github.com/gin-gonic/gin"
	"redenvelop-Prac/dal"
)

func main() {
	dal.InitDB()
	r := gin.Default()
	register(r)
	_ = r.Run() //listen and serve on 0.0.0.0:0000
}
