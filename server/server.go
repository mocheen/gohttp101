package server

import (
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	db *gorm.DB
	http.Server
	mux http.ServeMux
}

func (s *Server) Start() error {

	s.Server = http.Server{
		Addr:    ":8080",
		Handler: &s.mux,
	}
	return s.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.Shutdown(nil)
}

func (s *Server) RegisterRoute(f func(*Server)) {
	f(s)
}

func healthCheck(s *Server) {
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
}

func GetServer(db *gorm.DB) *Server {
	s := &Server{
		db: db,
	}
	s.RegisterRoute(healthCheck)
	return s
}
