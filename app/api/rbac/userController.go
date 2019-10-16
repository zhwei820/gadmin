package rbac

import (
	"github.com/gogf/gf/util/gconv"
	. "github.com/zhwei820/gadmin/app/api/base"
	"github.com/zhwei820/gadmin/app/service/service_model"
	"github.com/zhwei820/gadmin/utils/crypt"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
	"github.com/zhwei820/gadmin/app/api/api_model"
	"github.com/zhwei820/gadmin/app/model"
	"github.com/zhwei820/gadmin/app/service"
	"github.com/zhwei820/gadmin/utils/code"
	"time"
)

type UserController struct {
	BaseController
}

// @Summary user info
// @Description user info
// @Tags user
// @Success 200 {string} string	"ok"
// @router /rbac/user/info [get]
func (c *UserController) Info(r *ghttp.Request) {
	u := GetUser(r)
	if u != nil {
		Success(r, u.GetUserRoles())
		return
	}
	Fail(r, code.RESPONSE_ERROR, "获取用户信息失败")
}

// @Summary user list
// @Description user list
// @Tags user
// @Param	page	query 	integer	false		"page"
// @Param	page_size	query 	integer	false		"page_size"
// @Param	search	query 	string	false		"search"
// @Param	role_key	query 	string	false		"role_key"
// @Success 200 {string} string	"ok"
// @router /rbac/user [get]
func (c *UserController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	pageSize := r.GetInt("page_size", 10)
	roleKey := r.GetString("role_key", "")
	wheres := GetWhereFromRequest(r, nil, nil, []string{"username", "nickname", "email"})
	var userList struct {
		List  []service_model.GadminUserOut `json:"items"`
		Total int                           `json:"total"`
	}
	userList.List, userList.Total = service.GetPagedUser(wheres, roleKey, page, pageSize)
	Success(r, userList)
}

//
// @Summary CreateUser
// @Description CreateUser
// @Tags user
// @Param   CreateUser  body api_model.CreateUser true "CreateUser"
// @Success 200 {string} string	"ok"
// @router /rbac/user [post]
func (c *UserController) Post(r *ghttp.Request) {
        j, _ := r.GetJson()
	m := api_model.CreateUser{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	u, err := model.GetUserByName(m.Username)
	if err == nil && u.Id != 0 {
		Fail(r, code.RESPONSE_ERROR, "用户已存在")
		return
	}

	m.Password = crypt.EncryptPassword(m.Password)
	user := model.GadminUser{
		Username:   m.Username,
		Password:   m.Password,
		Nickname:   m.Nickname,
		Email:      m.Email,
		Phone:      m.Phone,
		AddUserId:  GetUserId(r),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	uid, err := user.Insert()
	if err != nil {
		glog.Debug(err.Error())
		glog.Debug(j.ToJsonString())
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	if uid > 0 {
		Success(r, "success")
	}

}

//
// @Summary UpdateUser
// @Description UpdateUser
// @Tags user
// @Param   UpdateUser  body api_model.UpdateUser true "UpdateUser"
// @Success 200 {string} string	"ok"
// @router /rbac/user [put]
func (c *UserController) Put(r *ghttp.Request) {
        j, _ := r.GetJson()
	m := api_model.UpdateUser{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String(), GetErrorMapForValid(e.Maps()))
		return
	}
	u, err := model.GetUserByName(m.Username)
	if u == nil || err != nil || u.Id == 0 {
		Fail(r, code.RESPONSE_ERROR, "用户不存在")
		return
	}
	umap := gconv.Map(m)
	delete(umap, "password")
	delete(umap, "passwordconfirm")

	if m.Password == "" {
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
		}
	} else {
		umap["password"] = crypt.EncryptPassword(m.Password)
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
		}
	}

	Success(r, "success")
}

//
// @Summary delete user
// @Description delete user
// @Tags user
// @Param	id	query 	integer	true		"id"
// @Success 200 {string} string	"ok"
// @router /rbac/user [delete]
func (c *UserController) Delete(r *ghttp.Request) {
	data, _ := r.GetJson()
	id := data.GetInt("id")
	if id < 1 {
		Fail(r, code.RESPONSE_ERROR, "不存在")
		return
	}
	u := new(model.GadminUser)
	user, err := u.GetById(id)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	if user.Username == model.ADMIN_NAME {
		Fail(r, code.RESPONSE_ERROR, "无权限删除管理员")
		return
	}
	res, _ := u.DeleteById(id)
	if res <= 0 {
		Fail(r, code.RESPONSE_ERROR, "删除失败")
		return
	}
	model.Enforcer.DeleteRolesForUser(user.Username)
	Success(r, "success")
}

//
// @Summary SetUserRole
// @Description SetUserRole
// @Tags role
// @Param   SetUserRole  body api_model.SetUserRole true "SetUserRole"
// @Success 200 {string} string	"ok"
// @router /rbac/role/userrole [put]
func (c *UserController) SetUserRole(r *ghttp.Request) {
        j, _ := r.GetJson()
	m := api_model.SetUserRole{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	for _, Username := range m.Usernames {
		service.SetUserRole(Username, m.RoleKeys)
	}
	Success(r, "success")
}
