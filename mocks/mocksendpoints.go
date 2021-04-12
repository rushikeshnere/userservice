package mocks 

import (
	"context"
	//"fmt"
	"github.com/go-kit/kit/endpoint"
	"userservice/models"
	"userservice/transport"
)

type Endpoints struct {
	CreateUserEndPoint       endpoint.Endpoint
}


func MakeCreateUserEndPoint(s Mocksservice) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CreateRequest)
		id, err := s.CreateUser(ctx, req.User)
		return transport.CreateResponse{Id: id, Error: err}, nil
	}
}

func MakeGetUserEndPoint(s Mocksservice) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetRequest)
		user, err := s.GetUser(ctx, req.Id)
		if(err != nil) {
			return transport.GetResponse{User: models.User{}, Error: err}, nil
		}
		return transport.GetResponse{User: *user, Error: err}, nil
	}
}

func MakeUpdateUserEndPoint(s Mocksservice) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.UpdateRequest)
		user, err := s.UpdateUser(ctx, req.User)
		if(err != nil) {
			return transport.UpdateResponse{User: models.User{}, Error: err}, nil
		}
		return transport.UpdateResponse{User: *user, Error: err}, nil
	}
}

func MakeDeleteUserEndPoint(s Mocksservice) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.DeleteRequest)
		result, err := s.DeleteUser(ctx, req.Id)
		return transport.DeleteResponse{Result: result, Error: err}, nil
	}
}