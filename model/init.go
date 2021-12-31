package model

import (
	"github.com/ismdeep/docker-hub-recorder/config"
	"github.com/ismdeep/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // load mysql driver
	"time"
)

var DB *gorm.DB

func init() {
	for {
		db, err := gorm.Open("mysql", config.Config.DSN)
		if err != nil {
			log.Warn("init", "info", "connect to db failed", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		DB = db
		break
	}

	DB.AutoMigrate(&Image{})
}
