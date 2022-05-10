package response

import "douyijn-simple/model"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FeedResponse struct {
	Response
	//omitempty 这则信息在原本的json数据如果没有，则不输出默认值
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
