package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPages(t *testing.T) {
	baseURL := "http://localhost:3000"

	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/5", 200},
		{"GET", "/articles/5/edit", 200},
		{"POST", "/articles/5", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 404},
	}

	for _, test := range tests {

		t.Logf("当前请求方式 %v 当前请求URL: %v \n", test.method, test.url)
		var (
			resp *http.Response
			err  error
		)
		switch {
		case test.method == "POST":
			data := make(map[string][]string)

			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}
		assert.NoError(t, err, "有错误发生, err 不为空")
		assert.Equal(t, test.expected, resp.StatusCode, "应返回状态码"+strconv.Itoa(test.expected))
	}

}
