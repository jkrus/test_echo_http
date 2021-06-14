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

	db, err := repository.NewJSONDB("users.json")
	if err != nil {
		logrus.Fatal(err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	serv := new(server.Server)

	err = serv.Run(viper.GetString("port"), handlers.InitRoutes())

	go func() {
		if err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
