package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	database "github.com/le0nar/time-control/db"
	"github.com/le0nar/time-control/internal/handler"
	"github.com/le0nar/time-control/internal/repository"
	"github.com/le0nar/time-control/internal/service"
	"github.com/le0nar/time-control/util"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	util.InitConfig()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := database.NewPostgresDB(database.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	port := viper.GetString("port")
	router :=  handler.InitRouter()
	router.Run("localhost:" + port)
}
