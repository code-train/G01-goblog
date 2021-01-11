package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/policies"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// ArticlesController struct
type ArticlesController struct {
	BaseController
}

// Show method
func (c *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouterVar("id", r)
	article, err := article.Get(id)

	if err != nil {
		c.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Article": article,
			"CanModifyArticle": policies.CanModifyArticle(article),
		}, "articles.show", "articles._article_meta")
	}
}

// Index 文章列表
func (c *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	var articles, pagerData, err = article.GetAll(r, 2)

	if err != nil {
		c.ResponseForSQLError(w, err)
	} else {

		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

// Create 创建页面
func (c *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")

}

// Store 创建文章
func (c *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	currentUser := auth.User()
	_article := article.Article{
		Title: r.PostFormValue("title"),
		Body:  r.PostFormValue("body"),
		UserID: currentUser.ID,
	}

	errors := requests.ValidateArticleForm(_article)

	if len(errors) == 0 {

		_article.Create()
		if _article.ID > 0 {
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
		c.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			c.ResponseForUnauthorized(w, r)
		} else {
			view.Render(w, view.D{
				"Article": _article,
				"Errors":  view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}

}

// Update method
func (c *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouterVar("id", r)

	_article, err := article.Get(id)

	if err != nil {
		c.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			c.ResponseForUnauthorized(w, r)
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
}

// Delete method
func (c *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouterVar("id", r)

	_article, err := article.Get(id)

	if err != nil {
		c.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			c.ResponseForUnauthorized(w, r)
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

}
