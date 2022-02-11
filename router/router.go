package router

import (
	"gin-pull-order/common"
	"gin-pull-order/module/cron/task/amazon"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {

	CronGroup := r.Group("/cron")
	{
		CronGroup.GET("/task/amazon/getorder", amazon.GetOrder)
		CronGroup.GET("/task/amazon/checkexception", amazon.CheckException)
	}

}

func Sign(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}
