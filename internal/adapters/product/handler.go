package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"market/internal/adapters"
	"market/internal/domain/product"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	solURL = "/products/:id"
	fewURL = "/products"
)

type handler struct {
	storage product.Storage
}

func NewHeandler(storage product.Storage) adapters.Handler {
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
	var dto product.CreateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input body"))
		return
	}
	r.Body.Close()

	br, err := h.storage.CreateRowDB(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(br)
}

func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto = product.DeleteProductDTO{Id: params.ByName("id")}

	if err := h.storage.DeleteRowDB(context.TODO(), dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Delete Product"))
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto product.UpdateProductDTO
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(br)
}

func (h *handler) GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto = product.GetProductDTO{Id: params.ByName("id")}
	br, err := h.storage.GetRowDB(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(br)
}

func (h *handler) ListItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println(r.Method, r.URL, r.TransferEncoding)

	list, err := h.storage.AllRowsDB(context.TODO(), params)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		fmt.Println(err)
	}
}
