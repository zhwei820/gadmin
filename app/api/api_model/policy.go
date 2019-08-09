package api_model

type SetPolicyByRole struct {
	Role    string   `json:"role" valid:"role      @required"`
	Policys []string `json:"policys" valid:"policys      @required"`
}

type UpdatePolicy struct {
	Path string `json:"path" valid:"path      @required"`
	Name string `json:"name" valid:"name      @required"`
}
