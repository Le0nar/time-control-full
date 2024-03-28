package main

import (
	"net/http"

	"github.com/Le0nar/time-control-full/tc-face-recognising/internal/handler"
	"github.com/Le0nar/time-control-full/tc-face-recognising/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	http.ListenAndServe(":3000", router)
}
