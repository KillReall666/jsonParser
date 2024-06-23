package storage

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"KillReall666/jsonParser.git/internal/config"
	"KillReall666/jsonParser.git/internal/model"

	jsoniter "github.com/json-iterator/go"
)

type Storage struct {
	repo map[string]model.PortData
	cfg  *config.Config
}

func New(cfg *config.Config) *Storage {
	return &Storage{
		repo: make(map[string]model.PortData),
		cfg:  cfg,
	}
}

func (s *Storage) LoadDataFromFile() error {
	file, err := os.Open(s.cfg.DataFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := jsoniter.NewDecoder(reader)

	for {
		if err = decoder.Decode(&s.repo); err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("decode err when load data from file: %v", err)
		}
	}

	return nil
}

func (s *Storage) SaveJSONData(name string, port model.PortData) {
	s.repo[name] = port
}

// UpdateJSONData This it an imitation of update data in storage
func (s *Storage) UpdateJSONData(name string, port model.PortData) {
	s.repo[name] = port
}

func (s *Storage) IsDataInStorage(port string) bool {
	_, ok := s.repo[port]
	return ok
}

func (s *Storage) PrintData() {
	fmt.Println(s.repo)
}
