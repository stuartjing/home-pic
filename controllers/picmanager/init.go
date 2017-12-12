package picmanager

import (
	"fmt"
	"home-pic/qiniu"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path"
	//	"strconv"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/admin/src/rbac"
	//	. "github.com/beego/admin/src"
	//"code.google.com/p/graphics-go/graphics"
	"home-pic/models"

	"github.com/BurntSushi/graphics-go/graphics"
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

func (c *InitController) Upload() {
	//接收上传的图片和描述  描述存在本地 图片本地不保留上传到七牛云上
	//	qiniu.UPloadFile()
	c.TplName = "easyui/picmanager/listok.tpl"

}

func (c *InitController) Save() {

	//接收上传的图片和描述  描述存在本地数据库  图片本地不保留上传到七牛云上 上传到七牛云成功后 讲返回的数据一起 保存到db中。配置好好域名，再做个图片列表页
	//	qiniu.UPloadFile()

	var pathstr string
	var name string
	var LocalP, qiniuimg string

	LocalP, _ = getCurrentPath()
	pic := c.GetStrings("pic[]")

	for i := 0; i < len(pic); i++ {

		pathstr = LocalP + pic[i]

		name = path.Base(pathstr)

		qiniuimg = qiniu.UPloadFile(pathstr, name)

		SaveImageDB("name", pic[i], qiniuimg, "描述图片，信息")

	}

	fmt.Println("-------------", pic)

	c.TplName = "easyui/picmanager/listok.tpl"

}

func (c *InitController) ShowUpload() {

	//这个页面是表单上传图片的页面
	//要支持批量上传
	c.TplName = "easyui/picmanager/showpuload.tpl"
}

func (c *InitController) Preview() {
	//预览图片
	var imgType, _ = c.GetInt("type", 0)
	var allfiles map[string][]*multipart.FileHeader = c.Ctx.Request.MultipartForm.File

	fmt.Println(allfiles, "wenjian=============")
	var keys []string
	var files []*multipart.FileHeader
	for k, vals := range allfiles {
		keys = append(keys, k)
		files = append(files, vals...)
	}
	var retArray []interface{}

	fmt.Println(files, "wenjian-------------")

	for i, h := range files {
		f, err := h.Open()
		fmt.Println(err, "wenjian-------------")
		defer f.Close()
		if err != nil {
			fmt.Println("api.ResponseUploadFileError", err)
			return
		}
		imageId, pathinfo, err := c.SaveFile(f, h, imgType)
		fmt.Println(imageId, err, "wenjian#################")
		if err != nil {
			fmt.Println("api.ResponseUploadFileError", err)
			return
		}
		var imageItem = map[string]interface{}{"image_id": imageId, "pathinfo": pathinfo, "file_name": h.Filename, "field": keys[i]}

		retArray = append(retArray, imageItem)

	}
	fmt.Println(retArray)

	//	var nodes interface{}
	c.Data["json"] = &map[string]interface{}{"total": 1, "rows": &retArray}
	c.ServeJSON()
	return
}

//上传图片,返回值为图片id
func (c *InitController) SaveFile(file multipart.File, fheader *multipart.FileHeader, imgType int) (int64, string, error) {

	var imageId int64 = 0
	str := "image/" + MakeDir()
	upload_path := beego.AppConfig.String("uploadpath")
	dir := upload_path + str
	//vpath := str + "/"
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return imageId, "", err
	}

	filename := fheader.Filename

	radomName := "tttt"
	extName := path.Ext(filename)
	largePath := dir + "/" + radomName + "_l" + extName
	//middlePath := dir + "/" + radomName + "_m" + extName
	//smallPath := dir + "/" + radomName + "_s" + extName

	dst, err := os.Create(largePath)
	defer dst.Close()
	if err != nil {
		return imageId, largePath, err
	}

	if _, err := io.Copy(dst, file); err != nil {
		return imageId, largePath, err
	}
	//imageSize, _ := beego.AppConfig.GetSection("image")
	//width_m, _ := strconv.Atoi(imageSize["width_m"])
	//height_m, _ := strconv.Atoi(imageSize["height_m"])

	//width_s, _ := strconv.Atoi(imageSize["width_s"])
	//height_s, _ := strconv.Atoi(imageSize["height_s"])

	//加载图片对象
	//	imgObj, err := ReadImage(largePath)

	if err != nil {
		return imageId, largePath, err
	}
	//err = ZoomImage(imgObj, middlePath, width_m, height_m)
	//if err != nil {
	//	return imageId, err
	//}
	//err = Thumbnail(imgObj, smallPath, width_s, height_s)

	//if err != nil {
	//	return imageId, err
	//} else {
	//l := vpath + radomName + "_l" + extName
	//m := vpath + radomName + "_m" + extName
	//s := vpath + radomName + "_s" + extName

	/*i := models.NewImage()
	rect := imgObj.Bounds()
	imgWidth := rect.Max.X
	imgHeight := rect.Max.Y
	imageId, err = i.InsertImage(filename, l, m, s, imgWidth, imgHeight, imgType)
	*/
	imageId = 4

	return imageId, largePath, nil
	//}
}

func MakeDir() string {
	return time.Now().Format("2006/01/02")
}

//等比缩放图片
func ZoomImage(src image.Image, savePath string, maxWidth int, maxHeight int) error {

	rect := src.Bounds()
	width := rect.Max.X
	height := rect.Max.Y

	_scale := float32(width) / float32(height)

	bw := width
	bh := height

	if bw > maxWidth {
		bw = maxWidth
		bh = int(float32(bw) / _scale)
	} else if bh > maxHeight {
		bh = maxHeight
		bw = int(float32(bh) * _scale)
	}

	dst := image.NewRGBA(image.Rect(0, 0, bw, bh))
	err := graphics.Scale(dst, src) //缩小图片
	if err != nil {
		return err
	}
	err = SaveImage(savePath, dst)

	return err
}

//裁剪图片
func Thumbnail(src image.Image, savePath string, maxWidth int, maxHeight int) error {

	rect := src.Bounds()
	width := rect.Max.X
	height := rect.Max.Y

	if width < maxWidth {
		maxWidth = width
	}
	if height < maxHeight {
		maxHeight = height
	}

	dst := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
	err := graphics.Thumbnail(dst, src) //缩小图片

	if err != nil {
		return err
	}
	err = SaveImage(savePath, dst)

	return err
}

//读取文件
func ReadImage(path string) (img image.Image, err error) {

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()
	//解决方案是上传base64  ?
	img, _, err = image.Decode(file) //解码图片

	return
}

//保存文件
func SaveImage(path string, img image.Image) (err error) {

	imgfile, err := os.Create(path)

	defer imgfile.Close()
	err = png.Encode(imgfile, img) //编码图片
	return
}

func SaveImageDB(name, pathinfo, qiniuimg, des string) {

	image := new(models.Images)

	image.Pathinfo = pathinfo
	image.Name = name
	image.Description = des
	image.Qiniuimg = qiniuimg

	a, err := models.AddImage(image)

	fmt.Println(a, err)
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
