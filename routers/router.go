package routers

import (
	"example.com/studyMongoDb/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.POST("/UserAdd", api.UserAdd)
	router.GET("/UserList", api.UserList)
	router.POST("/UserEditOne", api.UserEditOne)
	router.POST("/UserDelete", api.UserDelete)
	router.GET("/UserQuery/:key", func(c *gin.Context) {
		key := c.Param("key")
		api.UserQuery(c, key)
	})

	return router
}
