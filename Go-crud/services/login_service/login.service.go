package login_service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type LoginInformation struct {
	Username string
	Password string
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		Username: "blablabla@gmail.com",
		Password: "123456",
	}
}

func (info *LoginInformation) LoginUser(email string, password string) bool {
	return info.Username == email && info.Password == password
}
