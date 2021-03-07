package api

// TODO update customer model in the backend, CLI and UI
type Customer struct {
	Email    string `json:"email"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
}
