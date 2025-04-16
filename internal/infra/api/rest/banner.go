package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *Router) createBanner(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	imageURL := r.URL.Query().Get("imageURL")
	targetURL := r.URL.Query().Get("targetURL")
	text := r.URL.Query().Get("text")
	a.service.CreateBanner(r.Context(), title, text, imageURL, targetURL)
}

func (a *Router) getBanner(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	banner, err := a.service.GetBannerByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	resp, err := json.Marshal(banner)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Write(resp)
}

func (a *Router) updateBanner(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	title := r.URL.Query().Get("title")
	imageURL := r.URL.Query().Get("imageURL")
	targetURL := r.URL.Query().Get("targetURL")
	text := r.URL.Query().Get("text")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.service.UpdateBanner(r.Context(), id, title, text, imageURL, targetURL)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Router) deleteBanner(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.service.DeleteBanner(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Router) listBanners(w http.ResponseWriter, r *http.Request) {
	list := a.service.ListBanners(r.Context())

	resp, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
