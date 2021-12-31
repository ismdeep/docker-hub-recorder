package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/docker-hub-recorder/api"
	"github.com/ismdeep/docker-hub-recorder/config"
	"github.com/ismdeep/docker-hub-recorder/model"
)

func main() {
	router := gin.Default()
	router.GET("/api/images", api.GetImages) // get images
	fmt.Println(model.DB)
	if err := router.Run(config.Config.Bind); err != nil {
		panic(err)
	}
}
