package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/database"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
	"github.com/joho/godotenv"
)

const ( // FIXME
	defaultHTTPPort    = "8080"
	defaultDatabaseUrl = "postgres://postgres:12345678@127.0.0.1:5434/S7"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	httpAddr := net.JoinHostPort("127.0.0.1", getEnv("HTTP_PORT", defaultHTTPPort))
	dbConnStr := getEnv("DATABASE_URL", defaultDatabaseUrl)

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

func getEnv(name, defVal string) string {
	env := os.Getenv(name)
	if env == "" {
		return defVal
	}

	return env
}
