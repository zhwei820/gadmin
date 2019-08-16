package api

import (
	"fmt"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hailaz/gadmin/app/api/api_model"
	"github.com/hailaz/gadmin/app/service"
	"strings"

	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
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
// @Param	limit	query 	integer	false		"limit"
// @Success 200 {string} string	"ok"
// @router /policy [get]
func (c *PolicyController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)

	var list struct {
		List  []model.GadminPolicyconfig `json:"items"`
		Total int                        `json:"total"`
	}

	list.List, list.Total = service.GetPolicyList(page, limit, UNDEFIND_POLICY_NAME)

	Success(r, list)
}

//
// @Summary UpdatePolicy
// @Description UpdatePolicy
// @Tags policy
// @Param	name	query 	string	true		"name"
// @Param	policy	query 	string	true		"policy"
// @Success 200 {string} string	"ok"
// @router /policy [put]
func (c *PolicyController) Put(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.UpdatePolicy{}
	j.ToStruct(&m)

	if m.Name == UNDEFIND_POLICY_NAME {
		Fail(r, code.RESPONSE_ERROR)
	} else {
		err := service.UpdatePolicyByFullPath(m.Path, m.Name)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(r, "修改成功")
}

//
// @Summary GetPolicyByRole
// @Description GetPolicyByRole
// @Tags policy
// @Param	policy	query 	string	role		"role"
// @Success 200 {string} string	"ok"
// @router /policy/byrole [get]
func (c *PolicyController) GetPolicyByRole(r *ghttp.Request) {
	role := r.GetString("role")
	var list struct {
		List           []model.GadminPolicyconfig `json:"all_policy_items"`
		RolePolicyList []model.GadminPolicyconfig `json:"role_policy_items"`
		Total          int                        `json:"total"`
	}

	list.List, list.Total = service.GetPolicyList(1, -1, "")
	list.RolePolicyList = service.GetPolicyByRole(role)

	Success(r, list)
}

//
// @Summary SetPolicyByRole
// @Description SetPolicyByRole
// @Tags policy
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /policy [put]
func (c *PolicyController) SetPolicyByRole(r *ghttp.Request) {
	j := r.GetJson()
	m := api_model.SetPolicyByRole{}
	j.ToStruct(&m)

	var routerMap = make(map[string]model.RolePolicy)
	for _, item := range m.Policys {
		list := strings.Split(item, ":")
		path := list[0]
		atc := list[1]
		routerMap[fmt.Sprintf("%v %v %v", m.Role, path, atc)] = model.RolePolicy{Role: m.Role, Path: path, Atc: atc}
	}

	service.ReSetPolicy(m.Role, routerMap)

	Success(r, "success")
}
