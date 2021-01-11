package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// Article struct
type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
}

// Link method
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
