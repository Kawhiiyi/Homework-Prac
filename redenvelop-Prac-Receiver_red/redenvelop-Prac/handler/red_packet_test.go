package handler

import (
	"redenvelop-Prac/model"
	"testing"
)

func TestCheckParams(t *testing.T) {
	m := model.SendRpReq{}
	b := checkParams(m)
	t.Log(b)
}
