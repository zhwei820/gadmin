package api_model

type Login struct {
	Username string `valid:"user_name     @required"`
	Password string `valid:"password@required"`
	Kid      string `valid:"kid     @required"`
}
