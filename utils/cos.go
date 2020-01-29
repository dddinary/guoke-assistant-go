package utils

import (
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"guoke-helper-golang/config"
	"net/http"
	"net/url"
)

var CosCli *cos.Client

func init()  {
	bucket		:= config.CosConf.Bucket
	appId		:= config.CosConf.AppId
	region		:= config.CosConf.Region
	secretId	:= config.CosConf.SecretId
	secretKey	:= config.CosConf.SecretKey
	u, _ := url.Parse(fmt.Sprintf("http://%s-%s.cos.%s.myqcloud.com", bucket, appId, region))
	b := &cos.BaseURL{BucketURL:u}
	CosCli = cos.NewClient(b, &http.Client{
		Transport:     &cos.AuthorizationTransport{
			SecretID:     secretId,
			SecretKey:    secretKey,
		},
	})
}
