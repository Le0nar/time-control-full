package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/le0nar/time-control/internal/handler"
	"github.com/le0nar/time-control/internal/repository"
	"github.com/le0nar/time-control/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.DatabaseConfig{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	port := viper.GetString("port")
	router :=  handler.InitRouter()

	// TODO: move router.Run to goroutine
	router.Run("localhost:" + port)
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
}
