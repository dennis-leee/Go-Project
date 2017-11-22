package service

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

func home(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!") //这个写入到w的是输出到客户端
	}
}

func register(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.HTML(w, http.StatusOK, "register", struct{}{})
	}
}

func showResult(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() //获取用户信息
		//填充信息
		formatter.HTML(w, http.StatusOK, "result", struct {
			Name  string `json:"username"`
			ID    string `json:"studentID"`
			Phone string `json:"phone"`
			Mail  string `json:"mail"`
		}{Name: r.Form.Get("username"),
			ID:    r.Form.Get("studentID"),
			Phone: r.Form.Get("phone"),
			Mail:  r.Form.Get("mail")})
	}
}

func notImplemented(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "501 Not Implemented") //未实现
	}
}
