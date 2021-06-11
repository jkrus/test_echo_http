package main

import (
	"github.com/jkrus/test_echo_http/internal/service"
	"github.com/jkrus/test_echo_http/pkg/handler"
	"github.com/jkrus/test_echo_http/pkg/repository"
	"github.com/jkrus/test_echo_http/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewJSONDB("data.json")
	if err != nil {
		logrus.Fatal(err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	serv := new(server.Server)

	go func() {
		if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
