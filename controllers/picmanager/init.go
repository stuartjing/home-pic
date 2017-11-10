package picmanager

import (
	"fmt"

	//"github.com/astaxie/beego"
	"github.com/beego/admin/src/rbac"
	//	. "github.com/beego/admin/src"
)

type InitController struct {
	rbac.CommonController
	//beego.Controller
}

func (c *InitController) List() {
	//	this.Data["json"] = &map[string]interface{}{"total": 4, "rows": "wechat!"}
	//	this.ServeJSON()
	//	return
	fmt.Println("============================")

	c.TplName = "easyui/picmanager/list.tpl"

}
