package signatures

import (
	"github.com/gorilla/mux"
	"net/http"
)

type AppContext struct {
	Session *DatabaseSession
}

type AppContextHandler func(a *AppContext) http.HandlerFunc

type Route struct {
	Name           string
	Method         string
	Pattern        string
	AppHandlerFunc AppContextHandler
}

type Routes []Route

var routes = Routes{
	Route{
		"SignatureIndex",
		"GET",
		"/signatures",
		SignatureIndex,
	},
	Route{
		"SignatureCreate",
		"POST",
		"/signatures",
		SignatureCreate,
	},
	Route{
		"SignatureShow",
		"GET",
		"/signatures/{id}",
		SignatureShow,
	},
	Route{
		"SignatureDelete",
		"DELETE",
		"/signatures/{id}",
		SignatureDelete,
	},
}

func NewSignatureRouter(a *AppContext) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(route.AppHandlerFunc(a), route.Name))
	}

	return router
}
