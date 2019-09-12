package api_model

type CreateUser struct {
	Username        string `json:"username" valid:"username     @required"`
	Nickname        string `json:"nickname" valid:"nickname     @required"`
	Email           string `json:"email" valid:"email     @required|email"`
	Phone           string `json:"phone" valid:"phone     @required|phone"`
	Password        string `json:"password" valid:"password@required"`
	Passwordconfirm string `json:"passwordconfirm" valid:"passwordconfirm@required|same:password#|两次密码不一致，请重新输入"`
	Introduction string    `json:"introduction"`
	Avatar       string    `json:"avatar"`
}

type UpdateUser struct {
	Username        string `json:"username" valid:"username     @required"`
	Nickname        string `json:"nickname" valid:"nickname     @required"`
	Email           string `json:"email" valid:"email     @required|email"`
	Phone           string `json:"phone" valid:"phone     @required|phone"`
	Password        string `json:"password" valid:"password@required-with:passwordconfirm"`
	Passwordconfirm string `json:"passwordconfirm" valid:"passwordconfirm@same:password#两次密码不一致，请重新输入"`
	Introduction string    `json:"introduction"`
	Avatar       string    `json:"avatar"`
}

type GetUserParams struct {
	Username string `param:"username__contains"`
	Nickname string `param:"nickname__contains"`
	Email    string `param:"email__contains"`
	Search   string `param:"search"`
}
