package server

import (
	"github.com/willieso/baby-univ-biz-service/internal/routers"
	"github.com/willieso/baby-univ-biz-service/pkg/app"
	"github.com/willieso/baby-univ-biz-service/pkg/transport/http"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router
	// NOTE: register svc to http server

	return srv
}
