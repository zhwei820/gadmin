package api

import (
	"fmt"
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
// @Tags auth
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
// @Summary update menu
// @Description update menu
// @Tags user
// @Param   SignInInfo  body model.MenuOut true "MenuOut"
// @Success 200 {string} string	"ok"
// @router /menu [put]
func (c *PolicyController) Put() {
	data := c.Request.GetJson()
	name := data.GetString("name")
	path := data.GetString("policy")
	if name == UNDEFIND_POLICY_NAME {
		Fail(c.Request, code.RESPONSE_ERROR)
	} else {
		err := service.UpdatePolicyByFullPath(path, name)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}
	Success(c.Request, "修改成功")
}

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

func (c *PolicyController) SetPolicyByRole() {
	data := c.Request.GetJson()
	role := data.GetString("role")
	policys := data.GetStrings("policys")

	var routerMap = make(map[string]model.RolePolicy)
	for _, item := range policys {
		list := strings.Split(item, ":")
		path := list[0]
		atc := list[1]
		routerMap[fmt.Sprintf("%v %v %v", role, path, atc)] = model.RolePolicy{Role: role, Path: path, Atc: atc}
	}

	service.ReSetPolicy(role, routerMap)

	Success(c.Request, "success")
}
