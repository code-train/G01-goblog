package article

import (
	"goblog/pkg/route"
	"strconv"
)

// Article struct
type Article struct {
	ID    int64
	Title string
	Body  string
}

// Link method
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(a.ID, 10))
}
