package rbac

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
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
	pageSize := r.GetInt("page_size", 10)
	var list struct {
		List  []service_model.GadminRolePolicy `json:"items"`
		Total int                              `json:"total"`
	}
	list.List, list.Total = service.GetPagedRoleList(page, pageSize)

	Success(r, list)
}

// @Summary role list
// @Description role list
// @Tags role
// @Param	page	query 	integer	false		"page"
// @Param	page_size	query 	integer	false		"page_size"
// @Success 200 {string} string	"ok"
// @router /rbac/role [get]
func (c *RoleController) GetByRoleKey(r *ghttp.Request) {
	roleKey := r.GetString("role_key", "")
	ret, err := service.GetRoleByRolekey(roleKey)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	Success(r, ret)
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
	err := service.AddRole(m.RoleKey, m.Name)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	service.SetPolicyByRole(m.Policys, m.RoleKey)
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
		err := service.UpdateRoleByRoleKey(m.RoleKey, m.Name)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
		}
	}
	service.SetPolicyByRole(m.Policys, m.RoleKey)
	Success(r, "修改成功")
}

//
// @Summary delete role
// @Description delete role
// @Tags role
// @Param	role_key	query 	string	true		"role_key"
// @Success 200 {string} string	"ok"
// @router /rbac/role [delete]
func (c *RoleController) Delete(r *ghttp.Request) {
	roleKey := r.GetString("role_key")
	err := service.DeleteRole(roleKey)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	Success(r, "Delete")
}
