package model

// Addresses struct
type Address struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
}

// Response struct
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
