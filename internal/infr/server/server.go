package server

import (
	"context"
	companiesController "github.com/any/companies/internal/api/companies"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"sync"

	jwtController "github.com/any/companies/internal/api/jwt"
	"github.com/any/companies/internal/infr/logger"
)

type Server struct {
	httpServer *http.Server
	cfg        Config
}

func New(
	cfg Config,
	logger logger.Logger,
	jwtController jwtController.Controller,
	companiesController companiesController.Controller,
) (Server, error) {

	r := mux.NewRouter()

	/**** jwt ****/
	r.HandleFunc("/api/jwt/signin", jwtController.SigninHandler)
	r.HandleFunc("/api/jwt/refresh", jwtController.RefreshHandler)

	/**** companies ****/
	r.HandleFunc("/api/companies/{uuid}", companiesController.GetHandler).Methods(http.MethodGet)

	/**** protected companies ****/
	protected := r.Methods(http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()
	protected.Use(jwtController.JwtVerify)
	protected.HandleFunc("/api/companies", companiesController.CreateHandler).Methods(http.MethodPost)
	protected.HandleFunc("/api/companies/{uuid}", companiesController.PatchHandler).Methods(http.MethodPatch)
	protected.HandleFunc("/api/companies/{uuid}", companiesController.DeleteHandler).Methods(http.MethodDelete)

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
