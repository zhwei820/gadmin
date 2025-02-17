package options

import (
	"github.com/gogf/gf/net/ghttp"
	. "github.com/zhwei820/gadmin/app/api/base"
	"github.com/zhwei820/gadmin/app/service"
)

type OptionsController struct {
	BaseController
}
type InnerOption struct {
	// select组件等的选项
	Label string `json:"label" `
	Value string `json:"value" `
}
type OptionStruct struct {
	Type     string        `json:"type" `
	Model    string        `json:"model" `
	Label    string        `json:"label" `
	Options  []InnerOption `json:"options" `
	Multiple bool          `json:"multiple" `
}

// @Summary RoleOption
// @Description RoleOption
// @Tags options
// @Success 200 {string} string	"ok"
// @router /options/role-option [get]
func (c *OptionsController) RoleOption(r *ghttp.Request) {
	getRoleOptions := func() (innerOptions []InnerOption) {
		allRoleMap := service.GetAllRoleMap()
		for _, item := range allRoleMap {
			innerOptions = append(innerOptions, InnerOption{item.Name, item.RoleKey})
		}
		return
	}
	Success(r, []OptionStruct{
		{
			Type:     "select",
			Model:    "role_key",
			Label:    "角色名称",
			Options:  getRoleOptions(),
			Multiple: false,
		},
	})

}
