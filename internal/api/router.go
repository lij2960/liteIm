/************************************************************
 * Author:        jackey
 * Date:        2022/10/19
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package api

import (
	"liteIm/internal/api/controller"
	"net/http"
)

type RouterInterface interface {
	Get(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request))
	Post(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request))
}

type Router struct {
	exist bool
	RouterInterface
}

func (router *Router) Deal(w http.ResponseWriter, r *http.Request) {
	router.exist = false
	router.Post("/pushMsg", w, r, controller.PushMsg)

	// 用户操作
	router.Post("/user/register", w, r, controller.Register)
	router.Post("/user/remove", w, r, controller.Remove)

	// 用户组操作
	router.Post("/user/groupCreate", w, r, controller.GroupCreate)
	router.Post("/user/groupJoin", w, r, controller.GroupJoin)
	router.Post("/user/groupTransfer", w, r, controller.GroupTransfer)
	router.Post("/user/groupRemove", w, r, controller.GroupRemove)

	if !router.exist {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (router *Router) Get(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	if r.Method == http.MethodGet && r.URL.Path == path {
		router.exist = true
		handle(w, r)
	}
}

func (router *Router) Post(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	if r.Method == http.MethodPost && r.URL.Path == path {
		router.exist = true
		handle(w, r)
	}
}
