package qiniu

// 存储相关功能的引入包只有这两个，后面不再赘述
import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

//var localFile string = "/Users/stuartjing/a.jpeg"
var bucket string = "stuartjing"

//var key string = "github-x.png"
var accessKey string = "Hp0XfssJWGVk0csIzWEE3Pu6m_g1NO6x3KRMPYhA"
var secretKey string = "F-tcgPs7REDHkm9DUwLdpUgcl9xw3qJnu9iVgbvT"

func UPloadFile(localFile, key string) string {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
	//		Params: map[string]string{
	//			"x:name": "github logo",
	//		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(ret.Key, ret.Hash)
	return ret.Hash
}
