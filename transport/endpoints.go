package transport 

import (
	"context"
	//"fmt"
	"github.com/go-kit/kit/endpoint"
	"userservice/service"
	"userservice/models"
)

type Endpoints struct {
	CreateUserEndPoint       endpoint.Endpoint
}


func MakeCreateUserEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.CreateUser(ctx, req.User)
		return CreateResponse{Id: id, Error: err}, nil
	}
}

func MakeGetUserEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		user, err := s.GetUser(ctx, req.Id)
		if(err != nil) {
			return GetResponse{User: models.User{}, Error: err}, nil
		}
		return GetResponse{User: *user, Error: err}, nil
	}
}

func MakeUpdateUserEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		user, err := s.UpdateUser(ctx, req.User)
		if(err != nil) {
			return UpdateResponse{User: models.User{}, Error: err}, nil
		}
		return UpdateResponse{User: *user, Error: err}, nil
	}
}

func MakeDeleteUserEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		result, err := s.DeleteUser(ctx, req.Id)
		return DeleteResponse{Result: result, Error: err}, nil
	}
}