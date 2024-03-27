package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type service interface {
	RecogniseFace(file http.File) bool
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RecogniseFace(w http.ResponseWriter, r *http.Request) {
    // 10 << 20   10mb
	r.ParseMultipartForm(32 << 20) // limit your max input length!
    // in your case file would be fileupload
    file, header, err := r.FormFile("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // TODO: create folder for save images

    dst, err := os.Create(header.Filename)
    if err != nil {
        log.Println("error creating file", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "uploaded file")
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/recognise", h.RecogniseFace)

	return router
}