package controller

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/AttachmentModel"
	"main.go/app/v1/index/model/ProjectModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"os"
	"strings"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("up", up)
	route.Any("register")
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func up(c *gin.Context) {
	upload_file(c, false)
}

func up_full(c *gin.Context) {
	upload_file(c, true)
}

func upload_file(c *gin.Context, is_full bool) {
	token, ok := Input.Post("token", c, false)
	if !ok {
		return
	}
	proc := ProjectModel.Api_find_byToken(token)
	if len(proc) < 1 {
		RET.Fail(c, 404, nil, "未找到项目")
		return
	}
	file, ok := Input.Upload(c)
	if !ok {
		return
	}
	save_url := ""
	if is_full {
		save_url += proc["url"].(string) + "/"
	}
	file_info := AttachmentModel.Api_find_byMd5(file.Md5)
	if len(file_info) > 1 {
		RET.Success(c, 0, save_url+file_info["path"].(string), nil)
		return
	}

	if !strings.Contains(proc["ext"].(string), file.Ext) {
		RET.Fail(c, 400, nil, "上传的类型不在许可的范围内")
		return
	}
	if proc["size"].(int64) > file.Size {
		RET.Fail(c, 400, nil, "超出允许的上传大小")
		return

	}
	AttachmentModel.Api_insert(token, file.FileName, file.Path, file.Mime, file.Size, file.Ext, file.Md5)

	if proc["type"] == "local" || proc["type"] == "all" {
		save_url += file.Path
	}
	if proc["type"] == "oss" || proc["type"] == "all" {
		client, err := oss.New(proc["endpoint"].(string), proc["accesskey"].(string), proc["accesssecret"].(string))
		if err != nil {
			RET.Fail(c, 200, err.Error(), "OSS故障")
			return
		}
		bucket, err := client.Bucket(proc["bucket"].(string))
		if err != nil {
			RET.Fail(c, 200, err.Error(), "Bucket故障")
			return
		}
		err = bucket.PutObjectFromFile(file.FileName, file.Path)
		if err != nil {
			// HandleError(err)
		}
		if proc["type"] != "all" {
			os.Remove(file.Path)
		}
	}
}

func upload_base64(c *gin.Context) {

}
