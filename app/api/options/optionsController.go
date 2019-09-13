package options

import (
	"github.com/gogf/gf/g/net/ghttp"
	. "github.com/hailaz/gadmin/app/api/base"
	"github.com/hailaz/gadmin/app/service"
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

// @Summary user info
// @Description user info
// @Tags user
// @Success 200 {string} string	"ok"
// @router /options/role [get]
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
