package dto

type ResponseEmpty struct {
	Message string `json:"message"`
}

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}
