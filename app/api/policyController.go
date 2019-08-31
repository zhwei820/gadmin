package api

import (
	"fmt"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hailaz/gadmin/app/api/api_model"
	. "github.com/hailaz/gadmin/app/api/base"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/utils/code"
	"strings"
)

const (
	UNDEFIND_POLICY_NAME = "未命名"
)

type PolicyController struct {
	BaseController
}

// @Summary policy list
// @Description policy list
// @Tags policy
// @Param	page	query 	integer	false		"page"
// @Param	page_size	query 	integer	false		"page_size"
// @Success 200 {string} string	"ok"
// @router /rbac/policy [get]
func (c *PolicyController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	page_size := r.GetInt("page_size", 10)

	var list struct {
		List  []model.GadminPolicyconfig `json:"items"`
		Total int                        `json:"total"`
	}

	list.List, list.Total = service.GetPagedPolicyList(page, page_size)
	Success(r, list)
}

//
// @Summary UpdatePolicy
// @Description UpdatePolicy
// @Tags policy
// @Param   UpdatePolicy  body api_model.UpdatePolicy true "UpdatePolicy"
// @Success 200 {string} string	"ok"
// @router /rbac/policy [put]
func (c *PolicyController) Put(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.UpdatePolicy{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	if m.Name == UNDEFIND_POLICY_NAME {
		Fail(r, code.RESPONSE_ERROR, "不能设置为未命名")
		return
	} else {
		err := service.UpdatePolicyByFullPath(m.Path, m.Name, m.Label)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
		}
	}
	Success(r, "修改成功")
}

//
// @Summary GetPolicyByRole
// @Description GetPolicyByRole
// @Tags policy
// @Param	role	query 	string	role		"role"
// @Success 200 {string} string	"ok"
// @router /rbac/policy/byrole [get]
func (c *PolicyController) GetPolicyByRole(r *ghttp.Request) {
	role := r.GetString("role")
	var list struct {
		List           []model.GadminPolicyconfig `json:"all_policy_items"`
		RolePolicyList []model.GadminPolicyconfig `json:"role_policy_items"`
		Total          int                        `json:"total"`
	}

	list.List, list.Total = service.GetPagedPolicyList(1, 999999)
	list.RolePolicyList = service.GetPolicyByRole(role)

	Success(r, list)
}

//
// @Summary SetPolicyByRole
// @Description SetPolicyByRole
// @Tags policy
// @Param   SetPolicyByRole  body api_model.SetPolicyByRole true "SetPolicyByRole"
// @Success 200 {string} string	"ok"
// @router /rbac/policy/byrole [put]
func (c *PolicyController) SetPolicyByRole(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetPolicyByRole{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	var routerMap = make(map[string]model.RolePolicy)
	for _, item := range m.Policys {
		list := strings.Split(item, ":")
		path := list[0]
		act := list[1]
		routerMap[fmt.Sprintf("%v %v %v", m.Role, path, act)] = model.RolePolicy{Role: m.Role, Path: path, Act: act}
	}

	service.ReSetPolicy(m.Role, routerMap)

	Success(r, "success")
}
