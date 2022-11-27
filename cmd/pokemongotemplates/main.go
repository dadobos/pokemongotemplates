package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dadobos/pokemongotemplates/internal/infrastructure"
	"github.com/dadobos/pokemongotemplates/internal/infrastructure/app"
)

func ListenAndServeGoRoutine(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Println("Gracefully closing server..")
	}
}

func TerminationSignalWatcher(srv *http.Server) {
	// Make a channel to receive operating system signals.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
	log.Println("Shutting down server, because of received signal..")

	const GracefulShutdownRequestGraceSeconds = 10
	// The context is used to inform the server it has x seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(
		context.Background(),
		GracefulShutdownRequestGraceSeconds*time.Second,
	)

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	defer cancel()

	log.Println("Server exiting..")
}

func main() {
	http.DefaultClient.Timeout = 30 * time.Second

	infrastructure.ConfigSetup("config.toml", "./")

	srv := app.NewApplicationServer(nil)

	go ListenAndServeGoRoutine(srv.State.HTTPServer)
	TerminationSignalWatcher(srv.State.HTTPServer)

}
