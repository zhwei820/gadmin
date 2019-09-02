package service_model

type GadminRolePolicy struct {
	Id          int      `json:"id" `
	RoleKey     string   `json:"role_key" `
	Name        string   `json:"name" `
	Descrption  string   `json:"descrption" `
	PolicyKeys  []string `json:"policy_keys" `
	PolicyNames []string `json:"policy_names" `
}
