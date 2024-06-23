package service

import (
	"KillReall666/jsonParser.git/internal/model"
	"KillReall666/jsonParser.git/internal/storage"
)

type service struct {
	repo *storage.Storage
}

func New(r *storage.Storage) *service {
	return &service{
		repo: r,
	}
}

func (s *service) SaveJSONData(name string, port model.PortData) {
	s.repo.SaveJSONData(name, port)
}

func (s *service) UpdateJSONData(name string, port model.PortData) {
	s.repo.UpdateJSONData(name, port)
}

func (s *service) IsDataInStorage(port string) bool {
	return s.repo.IsDataInStorage(port)
}

func (s *service) PrintData() {
	s.repo.PrintData()
}
