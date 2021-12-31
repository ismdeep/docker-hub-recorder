package api

import "github.com/gin-gonic/gin"

func Fail(c *gin.Context, err error) {
	resp := make(map[string]interface{})
	resp["code"] = 1
	resp["msg"] = err.Error()
	c.JSON(200, resp)
}

func Success(c *gin.Context, data interface{}) {
	resp := make(map[string]interface{})
	resp["code"] = 0
	resp["data"] = data
	c.JSON(200, resp)
}

func Response(c *gin.Context, data interface{}, err error) {
	resp := make(map[string]interface{})
	if err != nil {
		resp["code"] = 1
		resp["msg"] = err.Error()
		c.JSON(200, resp)
		return
	} else {
		resp["code"] = 0
		resp["data"] = data
		c.JSON(200, resp)
		return
	}
}
