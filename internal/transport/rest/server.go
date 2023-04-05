package rest

import (
	"context"
	"log"
	"myapiproject/internal/config"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunHttp(cfg *config.Config, handler http.Handler) error {
	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	s.httpServer = &http.Server{
		Addr:              ":" + cfg.HTTPConfig.Port,
		Handler:           handler,
		ReadTimeout:       cfg.HTTPConfig.ReadTimeout,
		WriteTimeout:      cfg.HTTPConfig.WriteTimeout,
		MaxHeaderBytes:    cfg.HTTPConfig.MaxHeaderMegabytes << 20,
		ReadHeaderTimeout: cfg.HTTPConfig.ReadTimeout,
		IdleTimeout:       cfg.HTTPConfig.ReadTimeout,
		ErrorLog:          logger,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
