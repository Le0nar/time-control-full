package service

import "net/http"

type Serivce struct {}

func NewService() *Serivce {
	return &Serivce{}
}

func (s *Serivce) RecogniseFace(file http.File) bool  {
	return false
}
