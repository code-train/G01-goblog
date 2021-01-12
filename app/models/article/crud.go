package article

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文字
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)

	if err := model.DB.Preload("User").Preload("Category").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll method
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	viewData := _pager.Paging()

	var articles []Article
	_pager.Results(&articles)
	return articles, viewData, nil
}

// Create method
func (article *Article) Create() (err error) {
	fmt.Println("============ ARTICLE ============")
	fmt.Println(article)
	fmt.Println("============ ARTICLE ============")

	result := model.DB.Create(&article)

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

func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", cid), perPage)

	viewData := _pager.Paging()

	var articles []Article
	_pager.Results(&articles)
	return articles, viewData, nil
}

