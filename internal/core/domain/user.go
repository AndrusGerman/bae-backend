package domain

type User struct {
	Id      Id     `json:"Id"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
