package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUserName string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUserName: "pragmatic",
		authorizedPassword: "reviews",
	}
}

func (service *loginService) Login(username string, password string) bool {
	if service.authorizedUserName == username && service.authorizedPassword == password {
		return true
	}
	return false
}
