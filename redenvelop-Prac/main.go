package main //持久化存储
import (
	"github.com/gin-gonic/gin"
	"redenvelop-Prac/dal/db"
	"redenvelop-Prac/handler"
	"redenvelop-Prac/service"
)

func main() {
	db.InitDB()
	r := gin.Default()
	register(r)
	_ = r.Run() //listen and serve on 0.0.0.0:0000
}

func register(r *gin.Engine) {

	r.GET("/ping", handler.Demo)

	r.GET("/gin_demo/package_infos/:user_id", handler.QueryByUserId)

	r.POST("/gin_demo/package_infos", handler.InsertRecord)

	// 发放红包接口
	r.POST("/red-packet/send", service.SendRedPacket)
	// 查询发放记录
	r.GET("/red-packet/send/query", service.QuerySendRecords)
	// 领取红包接口
	r.POST("/red-packet/receive", service.ReceiveRedPacket)
	// 查询领取红包记录
	r.POST("/red-packet/receive/query", service.QueryReceiveRecords)
	// 查询发送红包记录分页
	r.POST("/red-packet/send/query/bypage", service.QuerySendRecordsByPage)
	// 查询领取红包记录分页
	r.POST("/red-packet/receive/query/bypage", service.QueryReceiveRecordsByPage)
	//导出发放记录
	r.POST("/red-packet/send/export", service.ExportSendRecords)
}
