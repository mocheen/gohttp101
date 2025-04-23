package server

import (
	"gohttp101/data"
	"net/http"
)

type Server struct {
	db data.Database
	http.Server
	mux http.ServeMux
}

func NewServer(db data.Database) *Server {
	return &Server{
		db: db,
	}
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

func GetServer(db data.Database) *Server {
	s := &Server{
		db: db,
	}
	s.RegisterRoute(healthCheck)
	return s
}
