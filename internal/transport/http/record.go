package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuchida-tamu/money-tracker-server/internal/record"
)

type RecordService interface{
	PostRecord(context.Context, record.Record)(record.Record, error)
	GetRecord(ctx context.Context, ID string)(record.Record, error)
	UpdateRecord(ctx context.Context, ID string, newRcd record.Record)(record.Record, error)
	DeleteRecord(ctx context.Context, ID string) error
}

type Response struct {
	message string
}

func (h *Handler) PostRecord(w http.ResponseWriter, r *http.Request){
	var rcd record.Record
	if err:= json.NewDecoder(r.Body).Decode(&rcd); err != nil {
		return
	}

	rcd, err:= h.Service.PostRecord(r.Context(), rcd)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(rcd); err != nil {
		panic(err)
	}
}

func (h *Handler) GetRecord(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rcd, err := h.Service.GetRecord(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(rcd); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateRecord(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rcd record.Record
	if err:= json.NewDecoder(r.Body).Decode(&rcd); err != nil {
		return
	}

	rcd, err := h.Service.UpdateRecord(r.Context(), id, rcd)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(rcd); err != nil {
		panic(err)
	}
}


func (h *Handler) DeleteRecord(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if id == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteRecord(r.Context(), id)
	if err != nil{
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{ message: "Successfully deleted" }); err != nil {
		panic(err)
	}
}