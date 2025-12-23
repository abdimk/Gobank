package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type APIServer struct {
	listenAddr string
	store Storage
}

func NewApiServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{listenAddr: listenAddr, store: store,}
}


func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handelAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handelGetAccountByID))

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
	account,err := s.store.GetAccounts()
	if err != nil{
		return WriteJSON(w, http.StatusBadRequest, err)
	}
	return WriteJSON(w, http.StatusOK, account)
	
}

func (s *APIServer) handelGetAccountByID(w http.ResponseWriter, r *http.Request) error{
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	

	return WriteJSON(w, http.StatusOK, &Account{})
}


func (s *APIServer) handelCreateAccount(w http.ResponseWriter, r *http.Request) error{
	CreateAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&CreateAccountReq); err != nil{
		return err
	}
	// fmt.Printf("%+v\n",CreateAccountReq)
	account := NewAccount(CreateAccountReq.FirstName, CreateAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, CreateAccountReq)
}


func (s *APIServer) handelDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}


func (s *APIServer) handelTransfer(w http.ResponseWriter, r *http.Request) error{
	return nil
}



func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


// function signature
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

