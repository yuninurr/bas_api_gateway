package usecase

type Login struct {
}

type LoginAuthenticator interface {
	Authenticate(username, password string) bool
}

func NewLogin() LoginAuthenticator {
	return &Login{}
}

func (a *Login) Authenticate(username, password string) bool {
	if username == "admin" && password == "password123" {
		return true
	}
	return false
}