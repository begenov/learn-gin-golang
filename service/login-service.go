package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	autorizedUsername string
	autorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		autorizedUsername: "pragmatic",
		autorizedPassword: "reviews",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.autorizedUsername == username && service.autorizedPassword == password
}
