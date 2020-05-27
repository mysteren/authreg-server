package apiserver

import (
	"net/http"

	"github.com/adam-hanna/jwt-auth/jwt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.devkeeper.com/authreg/server/internal/app/router"
	"gitlab.devkeeper.com/authreg/server/internal/app/store"
)

//
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	auth   *jwt.Auth
}

//
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//
func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting server")

	return http.ListenAndServeTLS(s.config.BindAddr, "cert.pem", "key.pem", router.New())
}

//
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

//
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	// s.store = st

	return nil
}
