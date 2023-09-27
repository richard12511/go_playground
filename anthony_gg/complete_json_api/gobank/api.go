package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{ listenAddr: listenAddr }
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", convertToHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", convertToHttpHandleFunc((s.handleGetAccount)))
	log.Println("JSON API server running and listening on port:", s.listenAddr)


	err := http.ListenAndServe(s.listenAddr, router)
	fmt.Println(err.Error())
}

func (s *APIServer) handleAccount(writer http.ResponseWriter, req *http.Request) error {
	if req.Method == "GET" {
		return s.handleGetAccount(writer, req)
	}
	if req.Method == "POST" {
		return s.handleCreateAccount(writer, req)
	}
	if req.Method == "DELETE" {
		return s.handleDeleteAccount(writer, req)
	}

	return fmt.Errorf("method not allowed")
}

func (s *APIServer) handleGetAccount(writer http.ResponseWriter, req *http.Request) error {
	account := NewAccount("Richard", "Schmidt")

	return writeJSON(writer, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(writer http.ResponseWriter, req *http.Request) error {
	return nil
}

type APIServer struct {
	listenAddr string
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func writeJSON(writer http.ResponseWriter, status int, val any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	return json.NewEncoder(writer).Encode(val)
}

func convertToHttpHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
			if err := fn(w, r); err != nil {
				writeJSON(w, http.StatusBadRequest, ApiError{ Error: err.Error()})
			}
		}
}

