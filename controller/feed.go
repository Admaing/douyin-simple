package controller

import (
	"douyijn-simple/model/response"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	c.JSON(200, response.FeedResponse{
		Response: response.Response{StatusCode:0},
		VideoList:
	})
}