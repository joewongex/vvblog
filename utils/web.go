package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonSuccess(c *gin.Context, datas ...gin.H) {
	data := gin.H{}
	if len(datas) > 0 {
		data = datas[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
