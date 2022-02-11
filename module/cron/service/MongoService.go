package service

import (
	"context"
	"fmt"
	"gin-pull-order/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//type Db interface {
//	GetConnection() *mongo.Client
//}

type MongoService struct {
	//connection MongoDb

}

//对象当做参数传入。内部使用变量mongoService即this
//方法的接收者只是另一个参数，因此应该相应地命名。 该名称不需要像方法参数那样具有描述性，因为它的作用是显而易见的并且没有任何记录目的。
func (mongoService *MongoService) getClient() *mongo.Client {
	uri := "mongodb://" + config.MONGO_USER + ":" + config.MONGO_PASSWORD + "@" + config.MONGO_HOST + ":" + config.MONGO_PORT + "/" + config.MONGO_AUTH_DB
	fmt.Println(uri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Fail Connected To MongoDB:", err)
	}
	// 检查连接情况
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Fail Connected To MongoDB!")
		log.Fatal(err)
	}
	return client
}

func (mongoService *MongoService) DbDcmBase() *mongo.Database {
	return mongoService.getClient().Database(config.MONGO_DB_DCM_BASE)
}
