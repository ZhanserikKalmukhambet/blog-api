package app

import (
	"github.com/ZhanserikKalmukhambet/blog-api/internal/config"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/handler"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/repository/pgrepo"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/service"
	"github.com/ZhanserikKalmukhambet/blog-api/pkg/http_server"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	db, err := pgrepo.New(
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)

	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}

	log.Println("connection success")

	srvs := service.New(db, cfg)
	hndlr := handler.New(srvs)

	server := http_server.New(
		hndlr.InitRouter(),
		http_server.WithPort(cfg.HTTP.Port),
		http_server.WithReadTimeout(cfg.HTTP.ReadTimeout),
		http_server.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		http_server.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil

}
