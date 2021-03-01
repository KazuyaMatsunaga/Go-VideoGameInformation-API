package server

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/db"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/controller"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) error {
	cs := db.NewDB(datasource)
	dbcon, err := cs.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = dbcon

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

	genreController := controller.NewGenre(s.db)
	r.Methods(http.MethodGet).Path("/genre/all").Handler(AppHandler{genreController.Index})

	return r
}