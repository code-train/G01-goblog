package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

// AuthController struct
type AuthController struct{}

// Register method
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth/register")
}

// DoRegister method
func (c *AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

}
