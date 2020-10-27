package main

import (
	"flag"
	"log"
	"os"

	server "github.com/Go-VideoGameInformation-API"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	s.Run(port)
}