package service

import (
	"context"
	"fmt"
	"gin-pull-order/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

//type Db interface {
//	GetConnection() *mongo.Client
//}

type MongoService struct {
	//connection MongoDb
}

//mongo连接单例，通过MongoService.getClient()获取
var mongoClient *mongo.Client

//加锁
var mutex sync.Mutex

//对象当做参数传入。内部使用变量mongoService即this
//方法的接收者只是另一个参数，因此应该相应地命名。 该名称不需要像方法参数那样具有描述性，因为它的作用是显而易见的并且没有任何记录目的。
func (mongoService *MongoService) getClient() *mongo.Client {
	if mongoClient == nil {
		var err error
		mutex.Lock()
		defer mutex.Unlock()
		uri := "mongodb://" + config.MONGO_USER + ":" + config.MONGO_PASSWORD + "@" + config.MONGO_HOST + ":" + config.MONGO_PORT + "/" + config.MONGO_AUTH_DB
		fmt.Println(uri)
		clientOptions := options.Client().ApplyURI(uri)
		//使用单例mongoClient，不能 用 := 号新定义
		mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Println("Fail Connected To MongoDB:", err)
		}
		// 检查连接情况
		err = mongoClient.Ping(context.TODO(), nil)
		if err != nil {
			fmt.Println("Fail Connected To MongoDB!")
			log.Fatal(err)
		}
	}
	return mongoClient
}

func (mongoService *MongoService) Disconnect() {
	err := mongoService.getClient().Disconnect(context.TODO())
	if err != nil {
		fmt.Println("mongo Disconnect Fail:", err)
	}
	mongoClient = nil
}

//获取数据库连接base
func (mongoService *MongoService) DbDcmBase() *mongo.Database {
	return mongoService.getClient().Database(config.MONGO_DB_DCM_BASE)
}

//获取数据库连接system
func (mongoService *MongoService) DbDcmSystem() *mongo.Database {
	return mongoService.getClient().Database(config.MONGO_DB_DCM_SYSTEM)
}
