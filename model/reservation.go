package model

type Reservation struct {
	Guest       Guest  `json:"guest"`
	Room        Room   `json:"room"`
	CheckInDate string `json:"check_in_date"`
	TotalPrice  int    `json:"total_price"`
	Payment     string `json:"payment"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
