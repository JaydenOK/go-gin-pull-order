package amazon

import (
	"context"
	"fmt"
	"gin-pull-order/module/cron/model"
	"gin-pull-order/module/cron/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOrder(c *gin.Context) {
	var task model.TaskModel
	mongoService := service.MongoService{}
	database := mongoService.DbDcmBase()
	//释放连接
	defer mongoService.Disconnect()

	//port := c.Query("port")
	//thread := c.DefaultQuery("thread", "10")
	//fmt.Println(port, thread)
	collection := database.Collection("task")
	//查询解析到task模型
	filter := bson.D{{"platform", "DMSAmazon"}}
	err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&task)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("单个查询: %+v\n", task)

	//2。查找多个
	var tasks []model.TaskModel
	filter = bson.D{{"platform", "DMSAmazon"}, {"status", 2}}
	//查询选项
	findOptions := options.Find()
	findOptions.SetLimit(5)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		fmt.Println(err)
	}
	for cursor.Next(context.TODO()) {
		var item model.TaskModel
		_ = cursor.Decode(&item)
		tasks = append(tasks, item)
	}

	fmt.Printf("多个查询: %+v\n", tasks)
	c.JSON(200, gin.H{
		"status": 0,
		"task":   tasks,
	})

}

func CheckException(c *gin.Context) {
	// 获取 Get 参数
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")

	c.JSON(200, gin.H{
		"v1":    "AddMember",
		"name":  name,
		"price": price,
	})
}
