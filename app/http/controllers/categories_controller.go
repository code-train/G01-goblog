package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (c CategoriesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "categories.create")
}

func (c CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)

	if len(errors) == 0 {
		_category.Create()
		if _category.ID > 0 {
			flash.Success("创建分类成功!")
			// indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(w, "创建文章分类失败")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors": errors,
		}, "categories.create")
	}
}

func (c CategoriesController) Show(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouterVar("id", r)
	_category, err := category.Get(id)

	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)

	if err != nil {
		c.ResponseForSQLError(w, err)
	}else {
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}


