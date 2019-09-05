package api_model

type PostRole struct {
	Name    string   `json:"name" valid:"name      @required"`
	RoleKey string   `json:"role_key" valid:"role_key      @required"`
	Policys []string `json:"policys" valid:"policys      @required"`
}

type SetUserRole struct {
	Usernames []string `json:"usernames" valid:"username      @required"`
	RoleKeys  []string `json:"role_keys" valid:"roles      @required"`
}
