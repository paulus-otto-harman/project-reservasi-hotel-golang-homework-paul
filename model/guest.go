package model

type Guest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}
