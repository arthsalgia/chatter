package models

type MessagesAll struct {
	ROWID          int    `gorm:"primaryKey;column:ROWID" json:"id"`
	AttributedBody []byte `gorm:"column:attributedBody" json:"-"`
	Ck_chat_id     string `json:"chat_id"`
	Is_from_me     int    `json:"is_from_me"`
	Date           int    `json:"date"`
}

func (MessagesAll) TableName() string {
	return "message"
}
