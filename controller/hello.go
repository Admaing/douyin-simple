package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "hello",
	})
}
