package service_model

import "github.com/hailaz/gadmin/app/model"

type GadminUserOut struct {
	model.GadminUser
	Roles []string `json:"roles" `
}
