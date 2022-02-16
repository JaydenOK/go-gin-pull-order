package model

type AccountListModel struct {
	Create_time string `json:"create_time"`
	Task_status int    `json:"task_status"`
	Update_time string `json:"update_time"`
	Platform    string `json:"platform"`
	Account_id  string `json:"account_id"`
	Type        int    `json:"type"`
	Short_name  string `json:"short_name"`
	Api_params  string `json:"api_params"`
}
