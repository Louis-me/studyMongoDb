package dbs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client     *mongo.Client     // database 话柄
	Collection *mongo.Collection // collection 话柄
)

// pool 连接池模式
func ConnectToDB(uri, name string) {

	cb := context.Background()
	// 设置连接超时时间
	ctx1, cancel := context.WithTimeout(cb, time.Duration(2))
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(50)
	// 发起链接
	client, err := mongo.Connect(ctx1, o)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断服务是不是可用
	if err = client.Ping(cb, readpref.Primary()); err != nil {
		log.Fatal(err)
		return
	}
	// 默认连接的数据库
	_db := client.Database(name)

	Client = client
	// 默认连接集合为users
	Collection = _db.Collection("users")
}
