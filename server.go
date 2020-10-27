package server

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	// "github.com/jmoiron/sqlx"
	"net/http"
	"os"
)

type Server struct {
	// db         *sqlx.DB
	router     *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	s.router = s.Route()
	return nil
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	r := mux.NewRouter()

	// Health check endpoint (http://localhost:1991/ping)
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	return r
}