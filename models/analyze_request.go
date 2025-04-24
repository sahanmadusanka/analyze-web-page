package models

type Request struct {
	Url string `json:"url" binding:"required" required:"$field is required"`
}
