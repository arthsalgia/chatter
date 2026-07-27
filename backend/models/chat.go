package models

type Chat struct {
	GroupID     string `json:"group_id"`
	DisplayName string `json:"display_name"`
}

func (Chat) TableName() string {
	return "chat"
}
