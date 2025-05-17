package http

import (
	netHttp "net/http"
)

type Response struct {
	Status int
	Data   interface{}
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ResponseOk(data any) Response {
	return Response{
		Status: netHttp.StatusOK,
		Data:   data,
	}
}

func ResponseCreated(data any) Response {
	return Response{
		Status: netHttp.StatusCreated,
		Data:   data,
	}
}

func ResponseNoContent() Response {
	return Response{
		Status: netHttp.StatusNoContent,
	}
}

func ResponseBadRequest() Response {
	return Response{
		Status: netHttp.StatusBadRequest,
		Data: ErrorResponse{
			Status:  netHttp.StatusBadRequest,
			Message: "Bad request",
		},
	}
}

func ResponseNotFound() Response {
	return Response{
		Status: netHttp.StatusNotFound,
		Data: ErrorResponse{
			Status:  netHttp.StatusNotFound,
			Message: "Not found",
		},
	}
}
