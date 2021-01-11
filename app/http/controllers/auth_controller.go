package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
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

	// 1. 初始化表单
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2.表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
			flash.Success("恭喜你注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal Server Error")
		}
	}

}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error": err.Error(),
			"Email": email,
			"Password": password,
		}, "auth.login")
	}

}

func (c AuthController) Logout(w http.ResponseWriter, r *http.Request)  {

	auth.Logout()
	flash.Success("您已成功退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}

