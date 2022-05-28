package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

type ServerConfig struct {
	Host string
	Port int
}

func NewServer(config *ServerConfig) *Server {
	router := mux.NewRouter()

	return &Server{
		&http.Server{
			Addr:    net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port)),
			Handler: router,
		},
	}
}

func (s *Server) StartServer() error {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return s.Shutdown(ctx)
}
