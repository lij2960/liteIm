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

func Router(w http.ResponseWriter, r *http.Request) {
	router := new(RouterFilter)
	router.Post("/pushMsg", w, r, controller.PushMsg)
}

type RouterInterface interface {
	Get(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request))
	Post(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request))
}

type RouterFilter struct {
	RouterInterface
}

func (rt *RouterFilter) Get(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodGet {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if r.URL.Path == path {
		handle(w, r)
	} else {
		http.Error(w, "not found", http.StatusNotFound)
	}
}

func (rt *RouterFilter) Post(path string, w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	if r.Method != http.MethodPost {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if r.URL.Path == path {
		handle(w, r)
	} else {
		http.Error(w, "not found", http.StatusNotFound)
	}
}
