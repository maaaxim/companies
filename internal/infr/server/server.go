package server

import (
	"context"
	"github.com/any/companies/internal/api"
	"github.com/any/companies/internal/infr/logger"
	"github.com/pkg/errors"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	cfg        Config
}

func New(
	cfg Config,
	logger logger.Logger,
	controller api.Controller,
) (Server, error) {
	r := mux.NewRouter()

	/**** nodes ****/
	r.HandleFunc(
		"/api/test",
		controller.SomeHandler,
	)

	//r.Use(api.JsonMiddleware)

	httpServer := &http.Server{
		Addr:         cfg.getHttpAddr(),
		Handler:      r,
		ReadTimeout:  cfg.Http.ReadTimeout,
		WriteTimeout: cfg.Http.WriteTimeout,
	}

	return Server{
		httpServer: httpServer,
		cfg:        cfg,
	}, nil
}

func (s Server) Listen() <-chan error {
	errChan := make(chan error)

	go func() {
		go func() {
			errChan <- s.httpServer.ListenAndServe()
		}()
	}()

	return errChan
}

func (s Server) Shutdown(ctx context.Context) error {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	errChan := make(chan error)
	go func() {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			errChan <- errors.Wrap(err, "shutdown http server error")
		}
		wg.Done()
	}()

	exitChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(exitChan)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err() //nolint:wrapcheck
	case err := <-errChan:
		return err
	case <-exitChan:
		return nil
	}
}
