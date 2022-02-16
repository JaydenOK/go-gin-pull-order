package model

type OrderModel struct {
	Order_id     string `json:"order_id"`
	Platform     string `json:"platform"`
	Account_id   string `json:"account_id"`
	Order_data   string `json:"order_data"`
	Order_time   string `json:"order_time"`
	Payment_time string `json:"payment_time"`
	Sync_flag    string `json:"sync_flag"`
	Sync_time    string `json:"sync_time"`
	Create_time  string `json:"create_time"`
	Update_time  string `json:"update_time"`
	Task_id      string `json:"task_id"`
	Exception    string `json:"exception"`
	Sync_times   string `json:"sync_times"`
}
