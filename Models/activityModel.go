package models

type Activity struct {
	ID                  int    `json:"id"`
	TaskID              int    `json:"task_id"`
	ActivityName        string `json:"activity_name"`
	ActivityDescription string `json:"activity_description"`
}
