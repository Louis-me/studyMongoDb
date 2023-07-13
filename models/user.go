package models

import (
	"context"
	"fmt"
	"log"

	"example.com/studyMongoDb/dbs"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

// // 序列化
// func (u *User) MarshalBinary() (data []byte, err error) {
// 	return json.Marshal(u)
// }

// // 反序列化
// func (u *User) UnmarshalBinary(data []byte) (err error) {
// 	return json.Unmarshal(data, u)
// }

func (u *User) GetList(ctx context.Context) (users []User, err1 error) {
	cur, err := dbs.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return
	}
	err = cur.All(ctx, &users)
	if err != nil {
		log.Fatal(err)
		return
	}
	cur.Close(ctx)
	return

	// log.Println("collection.Find curl.All: ", all)
	// for _, one := range all {
	// 	log.Println("Id:", one.Id, " - name:", one.Name, " - age:", one.Age)
	// }
}

func (u *User) Query() (user []User, err1 error) {
	cb := context.Background()
	fil := bson.M{"key": u.Key}
	// 模糊查询
	// bson.M{"Name": primitive.Regex{Pattern: u.name}}
	cur, err := dbs.Collection.Find(cb, fil)
	// dbs.Collection.FindOne(cb, fil).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return
	}
	err = cur.All(cb, &user)
	if err != nil {
		log.Fatal(err)
		return
	}
	cur.Close(cb)
	return

}

func (u *User) Add() (err error) {
	objId, err1 := dbs.Collection.InsertOne(context.TODO(), &u)
	if err1 != nil {
		fmt.Println("insert into error:", err)

	}
	fmt.Println("录入数据成功,", objId)
	return
}
func (u *User) EditOne() (err1 error) {
	fil := bson.M{"key": u.Key}
	update := bson.M{"$set": u}
	// .update({条件},{修改后的数据})
	updateResult, err := dbs.Collection.UpdateOne(context.Background(), fil, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
	return
}

func (u *User) Delete() (err1 error) {
	fil := bson.M{"key": u.Key}
	deleteResult, err := dbs.Collection.DeleteOne(context.Background(), fil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
	return
}
