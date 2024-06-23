package update

import (
	"encoding/json"
	"net/http"

	"KillReall666/jsonParser.git/internal/model"
)

type dataUpdater interface {
	SaveJSONData(name string, port model.PortData)
	UpdateJSONData(name string, port model.PortData)
	IsDataInStorage(port string) bool
	PrintData()
}

type Handler struct {
	dataUpdate dataUpdater
}

func NewUpdateHandler(dU dataUpdater) *Handler {
	return &Handler{
		dataUpdate: dU,
	}
}

// @Summary Update
// @Description Update information about ports in in-memory storage
// @Accept json
// @Produce json
// @Param Input body model.PortData true "port data"
// @Success 200 {integer} integer 1
// @Failure 400 {object} error
// @Router /update [post]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data map[string]model.PortData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "json format invalid", http.StatusBadRequest)
		return
	}

	var key string

	for k := range data {
		key = k
		break
	}

	if inStorage := h.dataUpdate.IsDataInStorage(key); !inStorage {
		h.dataUpdate.SaveJSONData(key, data[key])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("New data added successfully"))
		return
	} else {
		h.dataUpdate.UpdateJSONData(key, data[key])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data updated successfully"))
		return
	}
}
