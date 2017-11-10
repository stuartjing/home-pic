package routers

import (
	"home-pic/controllers"
	"home-pic/controllers/wechat"

	"home-pic/controllers/picmanager"

	"github.com/astaxie/beego"
	"github.com/beego/admin" //admin 包
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})

	beego.Router("/wechat/init/add", &wechat.InitController{}, "*:Add")

	//picmanager
	beego.Router("/picmanager/init/list", &picmanager.InitController{}, "*:List")

}
