package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Handler struct {
	Service CommService
	Router  *http.ServeMux
	Server  *http.Server
}

type CommService interface {
	EmpService
	UserService
}

func NewHandler(service CommService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = http.NewServeMux()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {

	h.Router.HandleFunc("POST /api/v1/employee", verifyJWT(h.CreateEmployee))
	h.Router.HandleFunc("GET /api/v1/employee/{id}", verifyJWT(h.GetEmployee))
	h.Router.HandleFunc("PUT /api/v1/employee", verifyJWT(h.UpdateEmployee))
	h.Router.HandleFunc("DELETE /api/v1/employee/{id}", verifyJWT(h.DeleteEmployee))

	h.Router.HandleFunc("POST /api/v1/user", h.CreateUser)
	h.Router.HandleFunc("POST /api/v1/signin", h.GetUser)
}

func (h *Handler) Serve() error {

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Server.Shutdown(context)
	return nil
}
