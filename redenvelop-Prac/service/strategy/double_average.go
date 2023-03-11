package strategy

import (
	_ "fmt"
	"math/rand"
	"time"
)

var min int64 = 1

func DoubleAverage(count, amount int64) int64 {

	if count == 1 {
		return amount
	}
	//计算出最大可用金额
	max := amount - min*count //计算出最大可用平均值
	avg := max / count        //二倍均值基础上再加上最小金额 防止出现金额为0
	avg2 := 2*avg + min       //随机红包金额序列元素，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}

func RandomAmount(totalAmount, totalNum int64, remainNum int64) int64 {
	if remainNum == 1 {
		return totalAmount
	}
	max := totalAmount - remainNum + 1
	amount := rand.Int63n(max) + 1
	return amount
}
