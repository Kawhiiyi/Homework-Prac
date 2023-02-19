package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"redenvelop-Prac/consts"
	"redenvelop-Prac/dal"
	"redenvelop-Prac/model"
	"redenvelop-Prac/utils"
	"time"
)

func Demo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func QueryByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	log.Printf("get userId from request  #{userId}")
	record, err := dal.QueryByUserId(userId)
	if err != nil {
		log.Printf("Can't find userId amount. userId : %v err : %v", userId, err)
		utils.RetErrorJson(c, consts.ParamsError)
	} else {
		rStr, _ := json.Marshal(record)
		utils.RetJsonWithData(c, string(rStr))
		return
	}
}

func InsertRecord(c *gin.Context) {
	var packageInfo model.PackageInfo
	err := c.BindJSON(&packageInfo)
	if err != nil {
		log.Printf("Bind package info error #{err}")
		utils.RetErrorJson(c, consts.ParamsError)
		return
	}
	record := &model.RpReceiveRecord{
		UserId:      packageInfo.UserId,
		GroupChatId: "insert_group002",
		RpId:        "insert_rp_002",
		Amount:      packageInfo.ReceiveAmount,
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
	}
	id, iErr := dal.InsertRecord(record)

	if iErr != nil {
		log.Printf("insert data err: #{err}\n")
		utils.RetErrorJson(c, consts.InsertError)
		return
	} else {
		log.Printf("i: #{id}\n")
		c.JSON(http.StatusOK, gin.H{
			"code":       "0",
			"message":    "success",
			"primary_id": id,
		})
	}
}
