package model

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}