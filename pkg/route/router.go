package route

import (
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var route *mux.Router

// SetRoute method
func SetRoute(r *mux.Router) {
	route = r
}

// Name2URL method
func Name2URL(routeName string, pairs ...string) string {

	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return url.String()
}

// GetRouterVar method
func GetRouterVar(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
