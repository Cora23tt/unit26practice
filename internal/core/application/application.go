package application

import "sync"

type Application struct {
	banners []Banner
	mux     *sync.Mutex
}

func NewApplication() *Application {
	return &Application{
		banners: make([]Banner, 0),
		mux:     &sync.Mutex{},
	}
}
