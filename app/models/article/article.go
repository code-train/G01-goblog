package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
)

// Article struct
type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User user.User `gorm:"save_associations:false"`
	CategoryID uint64 `gorm:"not null;default:4;index;" valid:"categoryID"`
	Category *category.Category
}

// Link method
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}

func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}