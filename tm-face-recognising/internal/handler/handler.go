package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type service interface {
	RecogniseFace(path string) bool
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

    const tempImagesDir = "_temp-images"

    if _, err := os.Stat(tempImagesDir); os.IsNotExist(err) {
        err := os.Mkdir(tempImagesDir, os.ModePerm)
        if err != nil{
            log.Fatal(err)
        }
    }

    pathToTempFile := tempImagesDir+ "/" + header.Filename

    createdFile, err := os.Create(pathToTempFile)
    if err != nil {
        log.Println("error creating file", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer createdFile.Close()
    
    if _, err := io.Copy(createdFile, file); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    isFaceRecognised := h.service.RecogniseFace(pathToTempFile)

    fmt.Fprint(w, strconv.FormatBool(isFaceRecognised))
}

// Initialization of router
func (h *Handler) InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/recognise", h.RecogniseFace)

	return router
}