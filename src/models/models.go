package models

type APIResponse struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}
