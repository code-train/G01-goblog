package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// Article struct
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link method
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
