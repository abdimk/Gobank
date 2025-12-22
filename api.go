package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}



type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r);  err != nil{
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}


type APIServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}


func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handelAccount))


	log.Println("API running on port:",s.listenAddr)


	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handelAccount(w http.ResponseWriter, r *http.Request) error{
	if r.Method == "GET"{
		return s.handelGetAccount(w, r)
	}

	if r.Method == "POST"{
		return s.handelCreateAccount(w,r)
	}

	if r.Method == "DELETE"{
		return s.handelDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handelGetAccount(w http.ResponseWriter, r *http.Request) error{
	// account := NewAccount("abdisa", "BK")
	// vars := mux.Vars(r)["id"]
	
	//db.get(id)

	return WriteJSON(w, http.StatusOK, &Account{})
}


func (s *APIServer) handelCreateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}


func (s *APIServer) handelDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}


func (s *APIServer) handelTransfer(w http.ResponseWriter, r *http.Request) error{
	return nil
}
