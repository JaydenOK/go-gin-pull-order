package model

//go json.Marshal时，结构体字段(首字母需要大写)，需要转换字段使用 `json:"platform"` 注释
//结构体属性、方法大小写是不同的，大写表示可以导出，引用到外部
//结构体不需要写全部字段信息，按需即可

type TaskModel struct {
	Platform          string `json:"platform"`
	Account_id        string `json:"account_id"`
	Status            int    `json:"status"`
	Create_begin_time string `json:"create_begin_time"`
	Create_end_time   string `json:"create_end_time"`
	Update_begin_time string `json:"update_begin_time"`
	Update_end_time   string `json:"update_end_time"`
	Execute_time      string `json:"execute_time"`
	Request_content   string `json:"request_content"`
	Response_content  string `json:"response_content"`
	Response_time     string `json:"response_time"`
	Type              int    `json:"type"`
	Message           string `json:"message"`
	Create_time       string `json:"create_time"`
	Update_time       string `json:"update_time"`
}
