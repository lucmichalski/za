package api

import (
	"net/http"

	"github.com/ncarlier/za/pkg/config"
	"github.com/ncarlier/za/pkg/middleware"
)

var commonMiddlewares = []middleware.Middleware{
	middleware.Gzip,
	middleware.Cors("*"),
	middleware.Logger("/healthz", "/badge"),
}

// NewRouter creates router with declared routes
func NewRouter(conf *config.Config) *http.ServeMux {
	router := http.NewServeMux()

	var middlewares = commonMiddlewares

	// Register HTTP routes...
	for _, route := range routes(conf) {
		handler := route.HandlerFunc(router, conf)
		for _, mw := range route.Middlewares {
			handler = mw(handler)
		}
		for _, mw := range middlewares {
			handler = mw(handler)
		}
		router.Handle(route.Path, handler)
	}

	return router
}
