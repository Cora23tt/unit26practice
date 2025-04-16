package rest

import (
	"banner/internal/core/application"
	"net/http"
)

type Router struct {
	mux     *http.ServeMux
	service *application.Application
}

func NewRouter(mux *http.ServeMux, service *application.Application) *Router {
	return &Router{
		mux:     mux,
		service: service,
	}

}

func (router *Router) Run() http.Handler {
	router.mux.HandleFunc("/", router.health)
	router.mux.HandleFunc("/banner/add", router.createBanner)
	router.mux.HandleFunc("/banner/get", router.getBanner)
	router.mux.HandleFunc("/banner/edit", router.updateBanner)
	router.mux.HandleFunc("/banner/delete", router.deleteBanner)
	router.mux.HandleFunc("/banners", router.listBanners)
	return router.mux
}
