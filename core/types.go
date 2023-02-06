package core

import "github.com/eliasacevedo/golang-microservice-template/server"

type Module interface {
	server.Route
}

type ErrorCode uint

type BaseResponse struct {
	Data      interface{} `json:"data"`
	ErrorCode ErrorCode   `json:"error_code"`
}

type PaginationParams struct {
	Page     uint `json:"page" form:"page"`
	Quantity uint `json:"quantity" form:"quantity"`
}

func NewPaginationParam(page uint, quantity uint) PaginationParams {
	if page == 0 {
		page = 1
	}

	if quantity == 0 {
		quantity = 1
	}

	return PaginationParams{
		Page:     page,
		Quantity: quantity,
	}
}
