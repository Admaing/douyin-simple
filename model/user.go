package model

type User struct {
	id            int64  `gorm:"primary_key:auto_increment" json:"id,omitempty"`
	Name          string `gorm:"type:varchar(200)" json:"name,omitempty"`
	FollowCount   int64  `gorm:"type:int" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"type:int" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"type:bool" json:"is_follow,omitempty"`
}
