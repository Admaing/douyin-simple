package model

type Video struct {
	//gorm.Model
	Id            int64  `gorm:"primary_key" json:"id,omitempty"`
	Author        User   `gorm:"embedded" json:"author"`
	PlayUrl       string `gorm:"type:varchar(200)" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"type:varchar(200)" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"type:varchar(200)" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"type:varchar(200)" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"type:varchar(200)" json:"is_favorite,omitempty"`
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
