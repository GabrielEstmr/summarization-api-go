package main_configurations_router

import (
	"sync"

	"github.com/gorilla/mux"
)

var Router *mux.Router = nil

var once sync.Once

func GetRouterBean() *mux.Router {
	once.Do(func() {

		if Router == nil {
			Router = getMuxRouterRouter()
		}

	})
	return Router
}

func getMuxRouterRouter() *mux.Router {
	return mux.NewRouter()
}
