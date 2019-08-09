package api

import (
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/api/api_model"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/library/code"
)

type RoleController struct {
	BaseController
}

// @Summary role list
// @Description role list
// @Tags role
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Param	username	query 	string	true		"username"
// @Success 200 {string} string	"ok"
// @router /role [get]
func (c *RoleController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	username := c.Request.GetString("username")
	var list struct {
		List         []model.GadminRoleconfig `json:"items"`
		UserRoleList []model.GadminRoleconfig `json:"role_items"`
		Total        int                      `json:"total"`
	}
	list.List, list.Total = service.GetRoleList(page, limit, UNDEFIND_POLICY_NAME)
	if username != "" {
		list.UserRoleList = service.GetRoleByUserName(username)
	}

	Success(c.Request, list)
}

//
// @Summary create role
// @Description create role
// @Tags role
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /policy [post]
func (c *RoleController) Post() {
	j := c.Request.GetJson()
	m := api_model.PostRole{}
	j.ToStruct(&m)

	err := service.AddRole(m.Role, m.Name)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}

	Success(c.Request, "Post")
}

//
// @Summary Update role
// @Description Update role
// @Tags role
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /role [put]
func (c *RoleController) Put() {
	j := c.Request.GetJson()
	m := api_model.PostRole{}
	j.ToStruct(&m)

	glog.Debug(m)
	if m.Name == UNDEFIND_POLICY_NAME {
		Fail(c.Request, code.RESPONSE_ERROR)
	} else {
		err := service.UpdateRoleByRoleKey(m.Role, m.Name)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(c.Request, "修改成功")
}

//
// @Summary delete role
// @Description delete role
// @Tags role
// @Param	role	query 	string	true		"role"
// @Success 200 {string} string	"ok"
// @router /role [delete]
func (c *RoleController) Delete() {
	role := c.Request.GetString("role")

	err := service.DeleteRole(role)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	Success(c.Request, "Delete")
}

//
// @Summary SetRoleByUserName
// @Description SetRoleByUserName
// @Tags role
// @Param   SetRoleByUserName  body api_model.SetRoleByUserName true "SetRoleByUserName"
// @Success 200 {string} string	"ok"
// @router /role/byuser [put]
func (c *RoleController) SetRoleByUserName() {
	j := c.Request.GetJson()
	m := api_model.SetRoleByUserName{}
	j.ToStruct(&m)

	service.SetRoleByUserName(m.Username, m.Roles)

	Success(c.Request, "success")
}

//
// @Summary SetRoleMenus
// @Description SetRoleMenus
// @Tags role
// @Param   SetRoleMenus  body api_model.SetRoleMenus true "SetRoleMenus"
// @Success 200 {string} string	"ok"
// @router /role/menu [put]
func (c *RoleController) SetRoleMenus() {
	j := c.Request.GetJson()
	m := api_model.SetRoleMenus{}
	j.ToStruct(&m)

	model.SetRoleMenus(m.Role, m.Menus)
	Success(c.Request, "success")
}
