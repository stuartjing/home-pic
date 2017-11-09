package routers

import (
	"wechatmanager/controllers"
	"wechatmanager/controllers/wechat"

	"github.com/astaxie/beego"
	"github.com/beego/admin" //admin 包
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})

	beego.Router("/wechat/init/add", &wechat.InitController{}, "*:Add")

}
