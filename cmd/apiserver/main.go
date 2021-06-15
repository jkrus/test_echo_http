package main

import (
	"context"
	"github.com/jkrus/test_echo_http/internal/service"
	"github.com/jkrus/test_echo_http/pkg/handler"
	"github.com/jkrus/test_echo_http/pkg/repository"
	"github.com/jkrus/test_echo_http/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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
	adrServ := viper.GetString("ip") + ":" + viper.GetString("port")
	go func() {
		if err := serv.Run(adrServ, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("test_echo_app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("test_echo_app stopped")

	if err := serv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server stopped: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
