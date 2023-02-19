package main

import (
	"github.com/gin-gonic/gin"
	"redenvelop-Prac/handler"
)

func register(r *gin.Engine) {
	r.GET("/ping", handler.Demo)

	r.GET("/gin_demo/package_infos/:user_id", handler.QueryByUserId)

	r.POST("/gin_demo/package_infos", handler.InsertRecord)

	r.POST("/red-packet/send", handler.SendRedPacket)

	r.GET("/red-packet/send/query", handler.QuerySendRecords)

	r.POST("/red-packet/receive", handler.ReceiveRedPacket)

	r.POST("/red-packet/receive/query", handler.QueryReceiveRecords)

}
