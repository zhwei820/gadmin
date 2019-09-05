package api_model

type CreateUser struct {
	Username        string `valid:"username     @required"`
	Nickname        string `valid:"nickname     @required"`
	Email           string `valid:"email     @required|email"`
	Phone           string `valid:"phone     @required|phone"`
	Password        string `valid:"password@required"`
	Passwordconfirm string `valid:"passwordconfirm@required|same:password#||两次密码不一致，请重新输入"`
}

type UpdateUser struct {
	Username        string `valid:"username     @required"`
	Nickname        string `valid:"nickname     @required"`
	Email           string `valid:"email     @required|email"`
	Phone           string `valid:"phone     @required|phone"`
	Password        string `valid:"password@"`
	Passwordconfirm string `valid:"passwordconfirm@same:password#||两次密码不一致，请重新输入"`
}

type GetUserParams struct {
	Username string `param:"username__contains"`
	Nickname string `param:"nickname__contains"`
	Email    string `param:"email__contains"`
	Search   string `param:"search"`
}
