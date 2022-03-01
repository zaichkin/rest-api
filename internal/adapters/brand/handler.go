package brand

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"market/internal/adapters"
	"market/internal/domain/brand"
	"net/http"
)

const (
	solURL = "/brands/:id"
	fewURL = "/brands"
)

type handler struct {
	storage brand.Storage
}

func NewHeandler(storage brand.Storage) adapters.Handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(solURL, h.GetItem)
	router.GET(fewURL, h.ListItems)
	router.POST(fewURL, h.AddItem)
	router.PUT(solURL, h.UpdateItem)
	router.DELETE(solURL, h.DeleteItem)
}

func (h *handler) AddItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto brand.CreateBrandDTO
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
	fmt.Println(string(buffer))
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
	var dto = brand.DeleteBrandDTO{Id: params.ByName("id")}

	if err := h.storage.DeleteRowDB(context.TODO(), dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Delete Brand"))
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto brand.UpdateBrandDTO
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
	var dto = brand.GetBrandDTO{Id: params.ByName("id")}
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
	br, err := h.storage.AllRowsDB(context.TODO())
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
