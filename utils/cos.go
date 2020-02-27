package utils

import (
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"guoke-assistant-go/config"
	"strconv"
	"time"
)

var cosClient *sts.Client

func init() {
	cosClient = sts.NewClient(
		config.CosConf.SecretId,
		config.CosConf.SecretKey,
		nil,
	)
}

func GetCosCredential(uid int) (map[string]interface{}, error) {
	region := config.CosConf.Region
	appid := config.CosConf.AppId
	bucket := config.CosConf.Bucket
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region: region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/user/" + strconv.Itoa(uid) + "/*",
					},
				},
			},
		},
	}
	creRes, err := cosClient.GetCredential(opt)
	if err != nil {
		return nil, err
	}
	res := map[string]interface{}{
		"tempSecretId":	creRes.Credentials.TmpSecretID,
		"tempSecretKey": creRes.Credentials.TmpSecretKey,
		"sessionToken":	creRes.Credentials.SessionToken,
		"expiredTime": creRes.ExpiredTime,
	}
	return res, nil
}
