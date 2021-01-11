package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/requests"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

// ArticlesController struct
type ArticlesController struct {
}

// Show method
func (c *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouterVar("id", r)
	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		view.Render(w, view.D{
			"Article": article,
		}, "articles.show")
	}
}

// Index 文章列表
func (c *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := article.GetAll()

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 Internal Server Error")
	} else {

		view.Render(w, view.D{
			"Articles": articles,
		}, "articles.index")
	}
}

// Create 创建页面
func (c *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {

	view.Render(w, view.D{}, "articles.create", "articles._form_field")

}

// Store 创建文章
func (c *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	_article := article.Article{
		Title: r.PostFormValue("title"),
		Body:  r.PostFormValue("body"),
	}

	errors := requests.ValidateArticleForm(_article)

	if len(errors) == 0 {

		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功, ID为"+_article.GetStringID())
			indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {
		data := view.D{
			"Article": _article,
			"Errors":  errors,
		}

		view.Render(w, data, "articles.create", "articles._form_field")
	}

}

// Edit method
func (c *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouterVar("id", r)

	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 NOT FOUND")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {

		view.Render(w, view.D{
			"Article": _article,
			"Errors":  view.D{},
		}, "articles.edit", "articles._form_field")
	}

}

// Update method
func (c *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouterVar("id", r)

	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not Found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {

		_article.Title = r.PostFormValue("title")
		_article.Body = r.PostFormValue("body")

		errors := requests.ValidateArticleForm(_article)

		fmt.Println("=======Article update ERRORS=======")
		fmt.Println(errors)
		if len(errors) == 0 {
			rowsAffected, err := _article.Update()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 Internal Server Error")
				return
			}

			if rowsAffected > 0 {
				showURL := route.Name2URL("articles.show", "id", id)
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "没有修改内容")
			}

		} else {
			data := view.D{
				"Article": _article,
				"Errors":  errors,
			}
			view.Render(w, data, "articles.edit", "articles._form_field")
		}

	}
}

// Delete method
func (c *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouterVar("id", r)

	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 NOT FOUND")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	} else {

		rowsAffected, err := _article.Delete()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		} else {
			if rowsAffected > 0 {
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 NOT FOUND")
			}
		}

	}

}
