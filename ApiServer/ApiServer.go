package ApiServer

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	router *mux.Router
	logger *logrus.Logger
	config *Config
}

func New() *ApiServer {
	return &ApiServer{
		router: mux.NewRouter(),
		logger: logrus.New(),
		config: NewConfig(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("start")
	return nil
}

func (s *ApiServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) ConfigureRouter() {

}
