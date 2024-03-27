package service

type Serivce struct {}

func NewService() *Serivce {
	return &Serivce{}
}

func (s *Serivce) RecogniseFace(path string) bool  {
	return false
}
