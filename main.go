package main

import (
	"fmt"
	"strings"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"github.com/gorilla/mux" 
	"context"
	"net/http"
	"encoding/json"
	transporthttp "github.com/go-kit/kit/transport/http" 
	
	"userservice/repository"
	"userservice/service"
	"userservice/transport"
)

func main() {

	dsn := "host=host.docker.internal user=postgres password=postgres dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err != nil) {
		fmt.Println("Error occurred while creating connection to database :", err)
	}	else {
		r := mux.NewRouter() 
		repository, _ := repository.New(db);

    	svc := service.NewService(repository)
  
 		CreateUserHandler := transporthttp.NewServer(  
			transport.MakeCreateUserEndPoint(svc),      
        	decodeCreateRequest,
			encodeResponse,
  		) 
  
		GetUserHandler := transporthttp.NewServer(
			transport.MakeGetUserEndPoint(svc),      
        	decodeGetRequest,
			encodeResponse,
		)

		UpdateUserHandler := transporthttp.NewServer(
			transport.MakeUpdateUserEndPoint(svc),
			decodeUpdateUserRequest,
			encodeResponse,      
		)

		DeleteUserHandler := transporthttp.NewServer(
			transport.MakeDeleteUserEndPoint(svc),
			decodeDeleteUserRequest,
			encodeResponse,      
		)
		
		r.Methods("POST").Path("/users").Handler(CreateUserHandler)
		r.Methods("GET").Path("/users/{id}").Handler(GetUserHandler)
		r.Methods("PUT").Path("/users").Handler(UpdateUserHandler)
		r.Methods("DELETE").Path("/users/{id}").Handler(DeleteUserHandler)
		fmt.Println("Starting server")
		http.ListenAndServe("0.0.0.0:8000", r) 		
	}
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	id := mux.Vars(r)["id"]
	var req = transport.GetRequest{Id: id}
	req.Id = id
	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.UpdateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
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
