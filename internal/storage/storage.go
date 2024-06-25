package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"KillReall666/jsonParser.git/internal/config"
	"KillReall666/jsonParser.git/internal/model"
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

	decoder := json.NewDecoder(reader)

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			log.Println("error reading token:", err)
			break
		}

		key, ok := token.(string)
		if !ok {
			continue // Skip token if not a string
		}

		var portData model.PortData
		err = decoder.Decode(&portData)
		if err != nil {
			log.Println("error decoding portDData:", err)
			break
		}

		s.repo[key] = portData
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
