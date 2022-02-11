package amazon

import (
	"context"
	"fmt"
	"gin-pull-order/module/cron/model"
	"gin-pull-order/module/cron/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOrder(c *gin.Context) {
	port := c.Query("port")
	thread := c.DefaultQuery("thread", "10")
	fmt.Println(port, thread)
	mongoService := service.MongoService{}
	database := mongoService.DbDcmBase()

	collection := database.Collection("task")
	filter := bson.D{{"platform", "DMSAmazon"}}
	cursor, err := collection.Find(
		context.Background(),
		filter,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	taskModel := model.TaskModel{}
	_ = cursor.Decode(&taskModel)

	fmt.Println("find task record:")
	fmt.Println(taskModel)

	// 断开客户端连接
	//err = client.Disconnect(context.TODO())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connection to MongoDB closed.")

	c.JSON(200, gin.H{
		"port":   port,
		"thread": thread,
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
