package auth

type Auth struct {
	Username   string
	Password   string
	Identifier string
}

func New(u, p, i string) Auth {
	a := Auth{
		Username:   u,
		Password:   p,
		Identifier: i,
	}
	return a
}

func (a *Auth) Login() bool {
	if len(a.Username) > 0 && len(a.Password) > 0 && len(a.Identifier) > 0 {
		return true
	}
	return false
}

func Check(accessToken string) bool {
	if len(accessToken) > 0 {
		return true
	}
	return false
}
