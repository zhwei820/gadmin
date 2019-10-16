package service_model

import "github.com/zhwei820/gadmin/app/model"

type GadminUserOut struct {
	model.GadminUser
	Roles []string `json:"roles" `
}
