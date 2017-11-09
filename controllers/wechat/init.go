package wechat

import (
	"fmt"

	"github.com/astaxie/beego"
	//	. "github.com/beego/admin/src"
)

type InitController struct {
	beego.Controller
}

func (c *InitController) Add() {
	fmt.Println("============================")
}
