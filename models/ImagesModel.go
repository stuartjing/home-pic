package models

import (
	"errors"
	"fmt"
	//	"log"
	"time"

	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	"github.com/astaxie/beego/validation"
	//	. "github.com/beego/admin/src/lib"
)

//图片表
type Images struct {
	Id          int64
	Name        string `orm:"size(32)" form:"Name" `
	Pathinfo    string `orm:"size(100)" `
	Qiniuimg    string `orm:"size(100)" `
	Addtime     int64  `orm:"size(11)" `
	Description string ` form:"Description" `
}

func (u *Images) TableName() string {
	return "images"
}

func init() {
	orm.RegisterModel(new(Images))
}

/************************************************************/

//get images list
func Getimageslist(page int64, page_size int64, sort string) (images []orm.Params, count int64) {
	o := orm.NewOrm()
	image := new(Images)
	qs := o.QueryTable(image)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&images)
	count, _ = qs.Count()
	return images, count
}

//添加 image
func AddImage(u *Images) (int64, error) {

	o := orm.NewOrm()
	image := new(Images)
	image.Name = u.Name
	image.Pathinfo = u.Pathinfo
	image.Qiniuimg = u.Qiniuimg
	image.Addtime = time.Now().Unix()
	image.Description = u.Description
	fmt.Println(image)
	id, err := o.Insert(image)
	return id, err
}

//更新image
func UpdateImage(u *Images) (int64, error) {

	o := orm.NewOrm()
	image := make(orm.Params)
	if len(u.Name) > 0 {
		image["Name"] = u.Name
	}
	if len(u.Pathinfo) > 0 {
		image["Pathinfo"] = u.Pathinfo
	}
	if len(u.Qiniuimg) > 0 {
		image["Qiniuimg"] = u.Qiniuimg
	}
	if len(u.Description) > 0 {
		image["Description"] = u.Description
	}

	if len(image) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Images
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(image)
	return num, err
}

func DelImageById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Images{Id: Id})
	return status, err
}
