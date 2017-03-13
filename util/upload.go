package util

import (
	"os"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	bucket    = os.Getenv("BUCKET")
	accesskey = os.Getenv("ACCESS_KEY")
	secretkey = os.Getenv("SECRET_KEY")
)

// PutRet ...
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

// Upload 上传到七牛
func Upload(fpath string) (*PutRet, error) {
	conf.ACCESS_KEY = accesskey
	conf.SECRET_KEY = secretkey

	c := kodo.New(0, nil)

	policy := &kodo.PutPolicy{
		Scope:   bucket,
		Expires: 3600,
	}

	token := c.MakeUptoken(policy)

	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	ret := &PutRet{}

	if err := uploader.PutFileWithoutKey(nil, ret, token, fpath, nil); err != nil {
		return nil, err
	}

	return ret, nil
}
