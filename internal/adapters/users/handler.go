package users

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"market/internal/adapters"
	"market/internal/domain/users"
	"market/pkg/usertools"
	"net/http"
)

const (
	solURL = "auth"
)

type handler struct {
	storage users.Storage
}

func NewHeandler(storage users.Storage) adapters.Handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {

	router.POST(solURL, h.SingIn)
	router.POST(solURL, h.SingUp)
}

func (h *handler) SingIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var sIn users.SingInDTO
	if err := json.NewDecoder(r.Body).Decode(&sIn); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	r.Body.Close()
	sIn.Password = usertools.HashPassword(sIn.Password)
	if err := h.storage.GetUser(context.TODO(), sIn); err != nil {
		log.Fatalln(err)
	}

}

func (h *handler) SingUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var sUp users.SingUpDTO
	if err := json.NewDecoder(r.Body).Decode(&sUp); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	r.Body.Close()
	sUp.Password = usertools.HashPassword(sUp.Password)
	if err := h.storage.CreateUser(context.TODO(), sUp); err != nil {
		log.Fatalln(err)
	}

}

func NewHandlers(storage users.Storage) adapters.Handler {
	return &handler{
		storage: storage,
	}
}
