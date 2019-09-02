package api_model

type SetPolicyByRole struct {
	Role    string   `json:"role" valid:"role      @required"`
	Policys []string `json:"policys" valid:"policys      @required"`
}

type UpdatePolicy struct {
	Path  string `json:"full_path" valid:"path      @required"`
	Name  string `json:"name" valid:"name      @required"`
	Label string `json:"label" valid:"label      @required"`
}
