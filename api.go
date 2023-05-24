package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status )
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w ).Encode(v)
}

type ApiError struct {
	Error string 
}
type apiFunc func(w http.ResponseWriter, http *http.Request) error  

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil{
			writeJSON(w, http.StatusBadRequest,ApiError{Error:err.Error()} )
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	log.Println("JSON api Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method  == "GET"{
		return s.handleGetAccount(w, r)
	}
	if r.Method  == "POST"{
		return s.handleCreateAccount(w, r)
	}
	if r.Method  == "DELETE"{
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, http *http.Request) error {
	return nil
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, http *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, http *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, http *http.Request) error {
	return nil
}