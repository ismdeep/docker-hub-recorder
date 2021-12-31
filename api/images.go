package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/docker-hub-recorder/model"
)

// GetImages get image list
// @Summary get image list
// @Author l.jiang.1024@gmail.com
// @Description get image list
// @Tags Image
// @Router	/api/v1/images [get]
func GetImages(c *gin.Context) {
	images := make([]*model.Image, 0)
	err := model.DB.Order("name asc").Find(&images).Error
	Response(c, images, err)
}
