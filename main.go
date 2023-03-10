package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/raymondgitonga/go-otel-sample/cmd/web"
	"github.com/raymondgitonga/go-otel-sample/internal/middleware/instrumentation"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}

	shutdown, tracer := instrumentation.SetupTracer(ctx)
	defer shutdown()

	app := web.NewApp(tracer)

	router, err := app.StartApp()
	if err != nil {
		log.Fatalf("error starting app: %s", err)
	}

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	fmt.Println(fmt.Sprintf("starting server on %s", port))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
