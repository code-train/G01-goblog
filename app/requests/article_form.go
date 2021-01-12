package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":	[]string{"required", "min:10"},
		"categoryID":[]string{"required"},
	}


	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min_cn:长度需大于 10",
		},
		"categoryID": []string{
			"required:请选择分类",
		},
	}

	options := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(options).ValidateStruct()
}
