package api

import (
	"fmt"
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
func (c *PolicyController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)

	var list struct {
		List  []model.GadminPolicyconfig `json:"items"`
		Total int                        `json:"total"`
	}

	list.List, list.Total = service.GetPolicyList(page, limit, UNDEFIND_POLICY_NAME)

	Success(c.Request, list)
}

//
// @Summary UpdatePolicy
// @Description UpdatePolicy
// @Tags policy
// @Param	name	query 	string	true		"name"
// @Param	policy	query 	string	true		"policy"
// @Success 200 {string} string	"ok"
// @router /policy [put]
func (c *PolicyController) Put() {
	j := c.Request.GetJson()
	m := api_model.UpdatePolicy{}
	j.ToStruct(&m)

	if m.Name == UNDEFIND_POLICY_NAME {
		Fail(c.Request, code.RESPONSE_ERROR)
	} else {
		err := service.UpdatePolicyByFullPath(m.Path, m.Name)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(c.Request, "修改成功")
}

//
// @Summary GetPolicyByRole
// @Description GetPolicyByRole
// @Tags policy
// @Param	role	query 	string	role		"role"
// @Success 200 {string} string	"ok"
// @router /policy/byrole [get]
func (c *PolicyController) GetPolicyByRole() {
	role := c.Request.GetString("role")
	var list struct {
		List           []model.GadminPolicyconfig `json:"all_policy_items"`
		RolePolicyList []model.GadminPolicyconfig `json:"role_policy_items"`
		Total          int                        `json:"total"`
	}

	list.List, list.Total = service.GetPolicyList(1, -1, "")
	list.RolePolicyList = service.GetPolicyByRole(role)

	Success(c.Request, list)
}

//
// @Summary SetPolicyByRole
// @Description SetPolicyByRole
// @Tags policy
// @Param   PostRole  body api_model.PostRole true "PostRole"
// @Success 200 {string} string	"ok"
// @router /policy [put]
func (c *PolicyController) SetPolicyByRole() {
	j := c.Request.GetJson()
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

	Success(c.Request, "success")
}
