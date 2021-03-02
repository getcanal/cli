package auth

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken string
type ProjectToken string

type ProjectCredentials struct {
	Namespace string `json:"namespace"`
}
