package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct{
	// https://zenn.dev/hsaki/articles/go-convert-json-struct
	Message string `json:"message"`
}
