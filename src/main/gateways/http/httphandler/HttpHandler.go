package main_gateways_http_httphandler

import (
	"github.com/gorilla/mux"
	main_gateways_http_controllers_beans "mpindicator/main/gateways/http/controllers/beans"
	"net/http"
	"sync"
)

type Routes struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

var once sync.Once

var RoutesBean *[]Routes = nil

func GetRoutesBean() *[]Routes {
	once.Do(func() {

		if RoutesBean == nil {
			RoutesBean = getFunctionBeans()
		}

	})
	return RoutesBean
}

func getFunctionBeans() *[]Routes {

	var RoutesConfig = []Routes{
		{
			URI:          "/users",
			Method:       http.MethodPost,
			Function:     main_gateways_http_controllers_beans.GetBeans().UserControllerBean.CreateUser,
			AuthRequired: false,
		},
		{
			URI:          "/indicators/reprocess",
			Method:       http.MethodPost,
			Function:     main_gateways_http_controllers_beans.GetBeans().IndicatorControllerBean.Reprocess,
			AuthRequired: false,
		},
		{
			URI:          "/indicators/{id}",
			Method:       http.MethodGet,
			Function:     main_gateways_http_controllers_beans.GetBeans().IndicatorControllerBean.FindIndicator,
			AuthRequired: false,
		},
	}

	return &RoutesConfig
}

func ConfigRoutes(r *mux.Router, routes []Routes) *mux.Router {
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
