package model

type Comment struct {
	Id         int64  `gorm:"primary_key" json:"id,omitempty"`
	User       User   `gorm:"embedded" json:"user"`
	UserID     int64  `gorm:"type:int;not null" json:"user_id,omitempty"`
	Content    string `gorm:"type:varchar(200)" json:"content,omitempty"`
	CreateDate string `gorm:"type:datetime" json:"create_date,omitempty"`
}
