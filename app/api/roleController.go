package api

import (
	"github.com/gogf/gf/g/net/ghttp"
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
func (c *RoleController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)
	username := r.GetString("username")
	var list struct {
		List         []model.GadminRoleconfig `json:"items"`
		UserRoleList []model.GadminRoleconfig `json:"role_items"`
		Total        int                      `json:"total"`
	}
	list.List, list.Total = service.GetRoleList(page, limit, UNDEFIND_POLICY_NAME)
	if username != "" {
		list.UserRoleList = service.GetRoleByUserName(username)
	}

	Success(r, list)
}

//
// @Summary create role
// @Description create role
// @Tags role
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /role [post]
func (c *RoleController) Post(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.PostRole{}
	j.ToStruct(&m)

	err := service.AddRole(m.Role, m.Name)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
	}

	Success(r, "Post")
}

//
// @Summary Update role
// @Description Update role
// @Tags role
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /role [put]
func (c *RoleController) Put(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.PostRole{}
	j.ToStruct(&m)

	glog.Debug(m)
	if m.Name == UNDEFIND_POLICY_NAME {
		Fail(r, code.RESPONSE_ERROR)
	} else {
		err := service.UpdateRoleByRoleKey(m.Role, m.Name)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(r, "修改成功")
}

//
// @Summary delete role
// @Description delete role
// @Tags role
// @Param	role	query 	string	true		"role"
// @Success 200 {string} string	"ok"
// @router /role [delete]
func (c *RoleController) Delete(r *ghttp.Request) {
	role := r.GetString("role")

	err := service.DeleteRole(role)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
	}
	Success(r, "Delete")
}

//
// @Summary SetRoleByUserName
// @Description SetRoleByUserName
// @Tags role
// @Param   SetRoleByUserName  body api_model.SetRoleByUserName true "SetRoleByUserName"
// @Success 200 {string} string	"ok"
// @router /role/byuser [put]
func (c *RoleController) SetRoleByUserName(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetRoleByUserName{}
	j.ToStruct(&m)

	service.SetRoleByUserName(m.Username, m.Roles)

	Success(r, "success")
}

//
// @Summary SetRoleMenus
// @Description SetRoleMenus
// @Tags role
// @Param   SetRoleMenus  body api_model.SetRoleMenus true "SetRoleMenus"
// @Success 200 {string} string	"ok"
// @router /role/menu [put]
func (c *RoleController) SetRoleMenus(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetRoleMenus{}
	j.ToStruct(&m)

	model.SetRoleMenus(m.Role, m.Menus)
	Success(r, "success")
}
