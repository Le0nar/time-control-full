package service

import (
	"log"

	"github.com/Kagami/go-face"
)

type Serivce struct {}

func NewService() *Serivce {
	return &Serivce{}
}

const modelsDir = "models"

func (s *Serivce) RecogniseFace(path string) bool  {

	// Init the recognizer.
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	faces, err := rec.RecognizeFile(path)
	if err != nil{
		log.Fatal(err)
	}

	hasFace := len(faces) > 0

// TODO: defer delete file

	return hasFace
}
