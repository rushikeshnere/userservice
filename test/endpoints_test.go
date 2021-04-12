package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	//"fmt"
	"strings"
	"bytes"
	"github.com/gorilla/mux" 
	transporthttp "github.com/go-kit/kit/transport/http" 
	"userservice/repository"
	"userservice/transport"
	"userservice/models"
	"userservice/mocks"
)

func TestGetUserWithCorrectId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/2", nil) 
	if err != nil {
		t.Fatal(err)
	}

	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	GetUserHandler := transporthttp.NewServer(
		mocks.MakeGetUserEndPoint(svc),      
		decodeGetRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("GET").Path("/users/{id}").Handler(GetUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserWithIncorrectId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/22", nil) 
	if err != nil {
		t.Fatal(err)
	}
	
	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	GetUserHandler := transporthttp.NewServer(
		mocks.MakeGetUserEndPoint(svc),      
		decodeGetRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("GET").Path("/users/{id}").Handler(GetUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


func TestCreateUserWithCorrectData(t *testing.T) {
	user := models.User{Id: "3", Name: "madhuri", Dob: "iw", Address: "P", Description: "D"}
	userData, e := json.Marshal(user)
	if(e != nil ){
		t.Fatal(e)
	}
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}
	
	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	CreateUserHandler := transporthttp.NewServer(
		mocks.MakeCreateUserEndPoint(svc),      
		decodeCreateRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("POST").Path("/users").Handler(CreateUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateUserWithInCorrectData(t *testing.T) {
	user := models.User{Id: "3", Name: "madhuri", Dob: "iw", Address: "P", Description: "D"}
	userData, e := json.Marshal(user)
	if(e != nil ){
		t.Fatal(e)
	}
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}

	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	CreateUserHandler := transporthttp.NewServer(
		mocks.MakeCreateUserEndPoint(svc),      
		decodeCreateRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("POST").Path("/users").Handler(CreateUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateUserWithCorrectId(t *testing.T) {
	user := models.User{Id: "1", Name: "modifiedname", Dob: "iw", Address: "P", Description: "D"}
	userData, e := json.Marshal(user)
	if(e != nil ){
		t.Fatal(e)
	}
	req, err := http.NewRequest("UPDATE", "/users", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}

	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	UpdateUserHandler := transporthttp.NewServer(
		mocks.MakeUpdateUserEndPoint(svc),      
		decodeUpdateRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("Update").Path("/users").Handler(UpdateUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateUserWithInCorrectId(t *testing.T) {
	user := models.User{Id: "5", Name: "modifiedname", Dob: "iw", Address: "P", Description: "D"}
	userData, e := json.Marshal(user)
	if(e != nil ){
		t.Fatal(e)
	}
	req, err := http.NewRequest("UPDATE", "/users", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}

	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	UpdateUserHandler := transporthttp.NewServer(
		mocks.MakeUpdateUserEndPoint(svc),      
		decodeUpdateRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("Update").Path("/users").Handler(UpdateUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteUserWithCorrectId(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	DeleteUserHandler := transporthttp.NewServer(
		mocks.MakeDeleteUserEndPoint(svc),      
		decodeDeleteRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("DELETE").Path("/users/{id}").Handler(DeleteUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


func TestDeleteUserWithInCorrectId(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/5", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	repository, _ := repository.New(nil);
	svc := mocks.NewService(repository)
	DeleteUserHandler := transporthttp.NewServer(
		mocks.MakeDeleteUserEndPoint(svc),      
		decodeDeleteRequest,
		encodeResponse,
	)

	router := mux.NewRouter() 
	router.Methods("DELETE").Path("/users/{id}").Handler(DeleteUserHandler)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)	

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	id := mux.Vars(r)["id"]
	var req = transport.GetRequest{Id: id}
	req.Id = id
	return req, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.UpdateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	id := mux.Vars(r)["id"]
	var req = transport.DeleteRequest{Id: id}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if resp, ok := response.(transport.GetResponse); ok && resp.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			result := &errorResponse{Message: getErrorMessageForError(resp.Error)}
			return json.NewEncoder(w).Encode(result)
	} 	

	if resp, ok := response.(transport.UpdateResponse); ok && resp.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := &errorResponse{Message: getErrorMessageForError(resp.Error)}
		return json.NewEncoder(w).Encode(result)
	} 
	
	if resp, ok := response.(transport.CreateResponse); ok && resp.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := &errorResponse{Message: getErrorMessageForError(resp.Error)}
		return json.NewEncoder(w).Encode(result)
	}
	
	if resp, ok := response.(transport.DeleteResponse); ok && resp.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		result := &errorResponse{Message: getErrorMessageForError(resp.Error)}
		return json.NewEncoder(w).Encode(result)
	}

	return json.NewEncoder(w).Encode(response)
}

func getErrorMessageForError(e error) string {
	if(strings.HasPrefix(e.Error(), "ERROR: duplicate key value violates unique constraint")) {
		return "User already exist"
	} else if(e.Error() == "record not found") {
		return "User not present"
	}
	return e.Error()
}

type errorResponse struct {
	Message string
}