package api_model

type PostRole struct {
	Name    string   `json:"name" valid:"name      @required"`
	RoleKey string   `json:"role_key" valid:"role_key      @required"`
	Policys []string `json:"policys" valid:"policys      @required"`
}

type SetRoleByUserName struct {
	Username string   `json:"username" valid:"username      @required"`
	Roles    []string `json:"roles" valid:"roles      @required"`
}
