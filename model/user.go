package model

type User struct {
	Id            int64  `gorm:"primaryKey" json:"id,omitempty"`
	Name          string `gorm:"type:varchar(200)" json:"name,omitempty"`
	FollowCount   int64  `gorm:"type:varchar(200)" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"type:varchar(200)" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"type:varchar(200)" json:"is_follow,omitempty"`
}
