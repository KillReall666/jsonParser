package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"

	_ "KillReall666/jsonParser.git/docs"
	"KillReall666/jsonParser.git/internal/config"
	"KillReall666/jsonParser.git/internal/handlers/update"
	"KillReall666/jsonParser.git/internal/service"
	"KillReall666/jsonParser.git/internal/storage"
)

// @title JSONParser App API
// @version 1.0
// @description API server for parsing json files that store information about ports.
// @termsOfService http://evil.com

// @contact.name API Support
// @contact.url http:///evil.com
// @contact.email codewarrior666@mail.ru

// @host localhost:8080
// @BasePath /

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	store := storage.New(cfg)

	err = store.LoadDataFromFile()
	if err != nil {
		panic(err)
	}
	serv := service.New(store)

	httpServer := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/update", update.NewUpdateHandler(serv).Update)
	http.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Starting server on", httpServer.Addr)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}
}
