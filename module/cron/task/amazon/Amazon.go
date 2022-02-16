package amazon

import (
	"context"
	"fmt"
	"gin-pull-order/module/cron/model"
	"gin-pull-order/module/cron/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetOrder(c *gin.Context) {
	mongoService := service.MongoService{}
	database := mongoService.DbDcmBase()
	//释放连接
	defer mongoService.Disconnect()
	getThread := c.Query("thread")
	thread := service.StringToInt(getThread)
	if thread < 0 {
		thread = 5
	}
	collection := database.Collection("task")
	filter := bson.D{{"platform", "DMSAmazon"}}

	//### 2 查找多个
	var tasks []model.TaskModel
	filter = bson.D{{"platform", "DMSAmazon"}, {"status", 0}}
	//查询选项
	findOptions := options.Find()
	findOptions.SetLimit(100)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		fmt.Println(err)
	}
	for cursor.Next(context.TODO()) {
		//查询数据解析到task结构体模型
		var item model.TaskModel
		_ = cursor.Decode(&item)
		tasks = append(tasks, item)
	}

	fmt.Printf("多个查询: %+v\n", tasks)
	//执行拉单
	pullOrderMain(tasks, thread)

	var response interface{}
	fmt.Println(response)

	c.JSON(200, gin.H{
		"status":   0,
		"response": response,
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

//任务主入口，启动多个协程拉单
func pullOrderMain(tasks []model.TaskModel, thread int) {
	//开启线程数
	//创建thread缓冲信道
	ch := make(chan model.TaskModel, thread)
	//遍历任务切片
	for index, task := range tasks {
		fmt.Println(index)
		//执行的任务推入有限缓冲信道，缓冲信道填满后，阻塞，有移出信道后，才能继续推入信道
		ch <- task
		go pullOrder(ch, task)
	}
	close(ch)
}

func pullOrder(ch chan model.TaskModel, task model.TaskModel) {
	pullPlatformOrderList(task)
	//执行完成，任务移出信道，接收新的任务
	<-ch
}

//拉取平台订单列表
func pullPlatformOrderList(task model.TaskModel) {
	url := "https://oms.goodcang.net/public_open/order/create_order"
	//randomNo := rand.Intn(999)
	httpService := service.HttpService{}
	response := httpService.Get(url)
	writeLog(task, response)
	//模拟延迟任务
	time.Sleep(2)
}

//拉取订单详情
func pullPlatformOrderDetail() {

}

func writeLog(task model.TaskModel, response string) {
	fmt.Println("[" + service.Date() + "]" + "platform:" + task.Platform + ",account_id:" + task.Account_id + ",response:" + response)
}
