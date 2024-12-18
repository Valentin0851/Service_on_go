package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"main.go/internal/handlers"
	"main.go/pkg/logging"
)

const (
	usersURl = "/users"
	userURl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURl, h.Getlist)
	router.GET(userURl, h.GetUserByUUID)
	router.POST(usersURl, h.CreateUser)
	router.PUT(userURl, h.UpdateUser)
	router.PATCH(userURl, h.ParticallyUpdateUser)
	router.DELETE(userURl, h.DeleteUser)
}

func (h *handler) Getlist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is getting user by uuid user"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is creating user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))
}

func (h *handler) ParticallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is dalete user"))
}
