package customer

type Customer struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
}

// New : Create a new customer
func New(Latitude string, Longitude string, UserID int, Name string) Customer {
	c := Customer{Latitude, Longitude, UserID, Name}
	return c
}
