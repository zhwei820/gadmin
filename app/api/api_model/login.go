package api_model

type Login struct {
	Username string `valid:"username     @required"`
	Password string `valid:"password@required"`
}
