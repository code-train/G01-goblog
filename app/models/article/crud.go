package article

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过 ID 获取文字
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll method
func GetAll() ([]Article, error) {
	var articles []Article

	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Create method
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	fmt.Println(result)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update method
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)

	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil

}

// Delete method
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}