package model

type Comment struct {
	Id         int64  `gorm:"primaryKey" json:"id,omitempty"`
	User       User   `gorm:"type:varchar(200)" json:"user"`
	Content    string `gorm:"type:varchar(200)" json:"content,omitempty"`
	CreateDate string `gorm:"type:varchar(200)" json:"create_date,omitempty"`
}
