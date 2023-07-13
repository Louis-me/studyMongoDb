package api

import (
	"context"
	"fmt"
	"net/http"

	"example.com/studyMongoDb/models"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		fmt.Println("data=", user)
		err := user.Add() // 调用model层的对应方法
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "新增失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "新增成功",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserList(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		users, err := user.GetList(context.Background()) // 调用model层的对应方法
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "获取列表失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":   "获取列表成功",
				"code":  1,
				"users": users,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserEditOne(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil {
		err := user.EditOne()
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "编辑失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "编辑成功",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

func UserDelete(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil {
		err := user.Delete()
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除成功",
				"code": 1,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}
func UserQuery(c *gin.Context, key string) {
	user := models.User{
		Key: key,
	}
	if c.Bind(&user) == nil {
		user, err := user.Query()
		if err != nil {

			c.JSON(http.StatusOK, gin.H{
				"msg":  "查询失败",
				"code": -1,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "查询成功",
				"code": 1,
				"user": user,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}
}
