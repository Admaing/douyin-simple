package model

type Video struct {
	//gorm.Model
	Id            int64  `gorm:"type:int;primary_key" json:"id,omitempty"`
	Author        User   `gorm:"embedded" json:"author"`
	PlayUrl       string `gorm:"type:varchar(200)" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"type:varchar(200)" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"type:int" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"type:int" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"bool" json:"is_favorite,omitempty"`
	UserID        int64  `gorm:"type:int;not null" json:"user_id,omitempty"`
}

//type VideoList string{
//
//}

//var err error

func GetVideoList(pageSize int, pageNum int) ([]Video, error) {
	var vedio []Video
	var total int64
	err = db.Preload("Video").Find(&vedio).Count(&total).Error
	if err != nil {
		return nil, err
	}
	return vedio, err
}
