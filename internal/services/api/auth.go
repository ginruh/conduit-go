package api

type AuthService interface {
	Login()
	Register()
}

type AuthServiceImpl struct{}

func (service AuthServiceImpl) Login() {

}

func (service AuthServiceImpl) Register() {

}
