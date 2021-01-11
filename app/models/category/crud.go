package category

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (c *Category) Create() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
