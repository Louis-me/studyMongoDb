package main

import (
	"context"

	"example.com/studyMongoDb/dbs"
	"example.com/studyMongoDb/routers"
)

func main() {

	dbs.ConnectToDB("mongodb://admin:123456@localhost:27017", "admin")
	router := routers.InitRouter()
	router.Run(":8000")
	// 优雅关闭连接
	defer dbs.Client.Disconnect(context.TODO())

}
