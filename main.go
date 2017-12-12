package main

import (
	_ "home-pic/routers"

	"github.com/astaxie/beego"
)

func main() {

	//定时清理本地图片-所有图片都存放到七牛云上

	//cleanPic()
	beego.Run()
}
