package api_model

type PostRole struct {
	Name string `json:"name" valid:"name      @required"`
	Role string `json:"role" valid:"role      @required"`
}

type SetRoleByUserName struct {
	Username string   `json:"username" valid:"username      @required"`
	Roles    []string `json:"roles" valid:"roles      @required"`
}

type SetRoleMenus struct {
	Role  string   `json:"role" valid:"role      @required"`
	Menus []string `json:"menus" valid:"roles      @required"`
}
