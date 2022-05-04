package model

type Client struct {
	ID     string `json:"client_id"`
	Secret string `json:"client_secret"`
}
