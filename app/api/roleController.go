package api

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hailaz/gadmin/app/api/api_model"
	. "github.com/hailaz/gadmin/app/api/base"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/app/service/service_model"
	"github.com/hailaz/gadmin/utils/code"
)

type RoleController struct {
	BaseController
}

// @Summary role list
// @Description role list
// @Tags role
// @Param	page	query 	integer	false		"page"
// @Param	page_size	query 	integer	false		"page_size"
// @Success 200 {string} string	"ok"
// @router /rbac/role [get]
func (c *RoleController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	page_size := r.GetInt("page_size", 10)
	var list struct {
		List  []service_model.GadminRolePolicy `json:"items"`
		Total int                              `json:"total"`
	}
	list.List, list.Total = service.GetPagedRoleList(page, page_size)

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
		return
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
		Fail(r, code.RESPONSE_ERROR, "不能设置为未命名")
		return
	} else {
		err := service.UpdateRoleByRoleKey(m.Role, m.Name)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
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
		return
	}
	Success(r, "Delete")
}

//
// @Summary SetRoleByUserName
// @Description SetRoleByUserName
// @Tags role
// @Param   SetRoleByUserName  body api_model.SetRoleByUserName true "SetRoleByUserName"
// @Success 200 {string} string	"ok"
// @router /rbac/role/userrole [put]
func (c *RoleController) SetRoleByUserName(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetRoleByUserName{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	service.SetRoleByUserName(m.Username, m.Roles)
	Success(r, "success")
}
