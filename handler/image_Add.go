package handler

import (
	"errors"
	"github.com/ismdeep/docker-hub-recorder/model"
)

func ImageAdd(name string) error {
	var cnt int64
	if err := model.DB.Model(&model.Image{}).Where("name=?", name).Count(&cnt).Error; err != nil {
		return err
	}

	if cnt >= 1 {
		return errors.New("already exists")
	}

	if err := model.DB.Create(&model.Image{
		Name: name,
	}).Error; err != nil {
		return err
	}

	return nil
}
