package transport 

import (
	"userservice/models"
)

type CreateRequest struct {
	User models.User
}

type CreateResponse struct {
	Id string
	Error error `json:"error,omitempty"`
}

type GetRequest struct {
	Id string
}

type DeleteRequest struct {
	Id string
}

type DeleteResponse struct {
	Result string
	Error error `json:"error,omitempty"`
}

type UpdateResponse struct {
	User models.User
	Error error `json:"error,omitempty"`
}

type UpdateRequest struct {
	User models.User
}

type GetResponse struct {
	User models.User
	Error error `json:"error,omitempty"`
}