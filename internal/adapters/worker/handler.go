package worker

import (
	"context"
	"encoding/json"
	"io"
	"market/internal/domain/worker"
	"net/http"

	"market/internal/adapters"

	"github.com/julienschmidt/httprouter"
)

const (
	solURL = "/workers/:id"
	fewURL = "/workers"
)

type handler struct {
	storage worker.Storage
}

func NewHeandler(storage worker.Storage) adapters.Handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(solURL, h.GetItem)
	router.GET(fewURL, h.ListItems)
	router.POST(fewURL, h.AddItem)
	router.PUT(solURL, h.UpdateItem)
	router.DELETE(fewURL, h.DeleteItem)
}

func (h *handler) AddItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto worker.CreateWorkerDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input body: " + err.Error()))
		return
	}
	r.Body.Close()

	br, err := h.storage.CreateRowDB(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(br)
}

func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto worker.DeleteWorkerDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input body"))
		return
	}
	r.Body.Close()

	if err := h.storage.DeleteRowDB(context.TODO(), dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Delete worker"))
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto worker.UpdateWorkerDTO
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input body"))
		return
	}
	if err := json.Unmarshal(buffer, &dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input body"))
		return
	}
	r.Body.Close()
	dto.Id = params.ByName("id")

	br, err := h.storage.UpdateRowDB(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(br)
}

func (h *handler) GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto = worker.GetWorkerDTO{Id: params.ByName("id")}
	br, err := h.storage.GetRowDB(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(br)
}

func (h *handler) ListItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	br, err := h.storage.AllRowsDB(context.TODO(), params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(br)
}
