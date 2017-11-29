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
	beego.Router("/picmanager/init/upload", &picmanager.InitController{}, "*:Upload")
	beego.Router("/picmanager/init/preview", &picmanager.InitController{}, "*:Preview")
	beego.Router("/picmanager/init/showupload", &picmanager.InitController{}, "*:ShowUpload")
	beego.Router("/picmanager/init/save", &picmanager.InitController{}, "*:Save")
}
