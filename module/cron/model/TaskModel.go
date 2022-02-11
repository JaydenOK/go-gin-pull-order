package model

type TaskModel struct {
	platform   string
	account_id string
	//Type              string `json:"type"`
	status            int
	update_start_time string
	update_end_time   string
	create_time       string
	message           string
}
