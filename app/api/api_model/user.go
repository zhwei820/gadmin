package api_model

type CreateUser struct {
	Username        string `valid:"user_name     @required"`
	Nickname        string `valid:"nick_name     @required"`
	Email           string `valid:"email     @required|email"`
	Phone           string `valid:"phone     @required|phone"`
	Password        string `valid:"password@required"`
	Passwordconfirm string `valid:"passwordconfirm@required|same:password#||两次密码不一致，请重新输入"`
}

type UpdateUser struct {
	Username        string `valid:"user_name     @required"`
	Nickname        string `valid:"nick_name     @required"`
	Email           string `valid:"email     @required|email"`
	Phone           string `valid:"phone     @required|phone"`
	Password        string `valid:"password@"`
	Passwordconfirm string `valid:"passwordconfirm@same:password#||两次密码不一致，请重新输入"`
}
