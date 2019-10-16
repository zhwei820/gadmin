package rbac

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"github.com/zhwei820/gadmin/app/api/api_model"
	. "github.com/zhwei820/gadmin/app/api/base"
	"github.com/zhwei820/gadmin/app/model"
	"github.com/zhwei820/gadmin/app/service"
	"github.com/zhwei820/gadmin/utils/code"
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
// @Param	search	query 	string	false		"search"
// @Success 200 {string} string	"ok"
// @router /rbac/policy [get]
func (c *PolicyController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	pageSize := r.GetInt("page_size", 10)
	search := r.GetString("search", "")

	var list struct {
		List  []model.GadminPolicyconfig `json:"items"`
		Total int                        `json:"total"`
	}

	list.List, list.Total = service.GetPagedPolicyList(search, page, pageSize)
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
// @Summary delete policy
// @Description delete policy
// @Tags policy
// @Param	policy	query 	string	true		"policy"
// @Success 200 {string} string	"ok"
// @router /rbac/policy [delete]
func (c *PolicyController) Delete(r *ghttp.Request) {
	path := r.GetString("path")

	err := service.DeletePolicy(path)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	Success(r, "Delete")
}

//
////
//// @Summary GetPolicyByRole
//// @Description GetPolicyByRole
//// @Tags policy
//// @Param	role	query 	string	role		"role"
//// @Success 200 {string} string	"ok"
//// @router /rbac/policy/byrole [get]
//func (c *PolicyController) GetPolicyByRole(r *ghttp.Request) {
//	role := r.GetString("role")
//	var list struct {
//		List           []model.GadminPolicyconfig `json:"all_policy_items"`
//		RolePolicyList []model.GadminPolicyconfig `json:"role_policy_items"`
//		Total          int                        `json:"total"`
//	}
//
//	list.List, list.Total = service.GetPagedPolicyList(1, 999999)
//	list.RolePolicyList = service.GetPolicyByRole(role)
//
//	Success(r, list)
//}
