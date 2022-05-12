package controller

import (
	"douyijn-simple/model"
	"douyijn-simple/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	videoList, err := model.GetVideoList(1, 10)
	if err != nil {
		c.JSON(200, response.Response{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("返回视频流错误", err),
		})
	}
	c.JSON(200, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: videoList,
	})
}
