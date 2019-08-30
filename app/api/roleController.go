package api

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hailaz/gadmin/app/api/api_model"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/utils/code"
)

type RoleController struct {
	BaseController
}

// @Summary role list
// @Description role list
// @Tags role
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Param	username	query 	string	false		"username"
// @Success 200 {string} string	"ok"
// @router /rbac/role [get]
func (c *RoleController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)
	username := r.GetString("username")
	var list struct {
		List         []model.GadminRoleconfig `json:"items"`
		UserRoleKeys []string                 `json:"user_role_keyss"`
		Total        int                      `json:"total"`
	}
	list.List, list.Total = service.GetPagedRoleList(page, limit)
	if username != "" {
		list.UserRoleKeys = service.GetRoleKeysByUserName(username)
	}

	Success(r, list)
}

//
// @Summary create role
// @Description create role
// @Tags role
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /rbac/role [post]
func (c *RoleController) Post(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.PostRole{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
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
// @router /rbac/role [put]
func (c *RoleController) Put(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.PostRole{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
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
// @router /rbac/role [delete]
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
// @router /rbac/role/byuser [put]
func (c *RoleController) SetRoleByUserName(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetRoleByUserName{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	_ = service.SetRoleByUserName(m.Username, m.Roles)
	Success(r, "success")
}
