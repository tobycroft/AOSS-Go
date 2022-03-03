package controller

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/v1/index/model/ProjectModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("login", loginss)
	route.Any("upload", upload)
	route.Any("register")
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func loginss(c *gin.Context) {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(0, json)
}

func upload(c *gin.Context) {
	file, ok := Input.PostFile(c)
	if !ok {
		return
	}
	fmt.Println(file)
}

func up(c *gin.Context) {
	client, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}
}

func upload_file(c *gin.Context) {
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

}

func upload_base64(c *gin.Context) {

}
