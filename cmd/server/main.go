package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/config"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/database"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
)

const (
	defHttpIp   = "127.0.0.1"
	defHttpPort = "8080"
	defDbUrl    = "postgres://postgres:postgres@127.0.0.1:5432/Api" // FIXME use default container conf in future
)

func main() {
	if config.ErrLoadingEnvs != nil {
		log.Fatalln("Error loading .env file")
	}

	httpAddr := net.JoinHostPort(config.GetEnv("SERVER_IP", defHttpIp), config.GetEnv("SERVER_PORT", defHttpPort))
	dbConnStr := config.GetEnv("DATABASE_URL", defDbUrl)

	dbc, err := database.Connect(dbConnStr)
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	env := &handler.Env{httpAddr, dbc}
	srv := app.NewServer(env)

	go srv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv.Shutdown(ctx)
	os.Exit(0)
}
