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

	"github.com/JC5LZiy3HVfV5ux/nord/internal/handlers"
	"github.com/JC5LZiy3HVfV5ux/nord/internal/repositories"
	"github.com/JC5LZiy3HVfV5ux/nord/internal/services"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

type ServerConfig struct {
	Host            string
	Port            int
	Key             string
	Lang            string
	Unit            string
	RedisHost       string
	RedisPort       int
	RedisPassword   string
	RedisDB         int
	RedisExpiration time.Duration
}

func NewServer(config *ServerConfig) (*Server, error) {
	router := mux.NewRouter()

	openweatherClient, err := setupOpenweatherClient(config)
	if err != nil {
		return nil, err
	}

	redisClient, err := setupRedisClient(config)
	if err != nil {
		return nil, err
	}

	repositories := repositories.NewRepository(redisClient, config.RedisExpiration)

	services := services.NewServices(openweatherClient, repositories)

	handlers.RegisterHandlers(router, services)

	return &Server{
		&http.Server{
			Addr:    net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port)),
			Handler: router,
		},
	}, nil
}

func (s *Server) StartServer() error {
	log.Println("Start server...")

	go func() {
		log.Printf("http://%s/api/v1/ping", s.Addr)

		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.Shutdown(ctx)
}

func setupOpenweatherClient(config *ServerConfig) (*openweather.Client, error) {
	openweatherClient, err := openweather.NewClient(config.Key, &http.Client{Timeout: time.Second * 5})
	if err != nil {
		return nil, err
	}

	if err := openweatherClient.SetLang(config.Lang); err != nil {
		log.Println(err)
	}

	if err := openweatherClient.SetUnit(config.Unit); err != nil {
		log.Println(err)
	}

	return openweatherClient, nil
}

func setupRedisClient(config *ServerConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(config.RedisHost, fmt.Sprintf("%d", config.RedisPort)),
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
