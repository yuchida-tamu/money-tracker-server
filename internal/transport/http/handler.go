package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type RecordService interface{}

type Handler struct {
	Router *mux.Router
	Service RecordService
	Server *http.Server
}

func NewHandler(service RecordService) *Handler{
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h* Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello world")
	})
}

func (h *Handler) Serve() error {
	// listen to requests in a non-blocking manner
	go func(){
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	// create a channel to listen to os.Interrup signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // blocking until the channel catch the signal

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// cancel will called 15seconds after
	defer cancel() 
	h.Server.Shutdown((ctx))

	log.Println("shut down gracefully")

	return nil
}