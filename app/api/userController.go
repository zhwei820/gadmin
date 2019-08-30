package api

import (
	. "github.com/hailaz/gadmin/app/api/base"

	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gvalid"
	"github.com/hailaz/gadmin/app/api/api_model"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/utils"
	"github.com/hailaz/gadmin/utils/code"
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
	}
	Fail(r, code.RESPONSE_ERROR, "获取用户信息失败")
}

// @Summary user list
// @Description user list
// @Tags user
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Param	search	query 	string	false		"search"
// @Success 200 {string} string	"ok"
// @router /rbac/user [get]
func (c *UserController) Get(r *ghttp.Request) {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)
	wheres := GetWhereFromRequest(r, nil, nil, []string{"user_name", "nick_name", "email"})
	var userList struct {
		List  []model.GadminUser `json:"items"`
		Total int                `json:"total"`
	}
	userList.List, userList.Total = service.GetPagedUser(wheres, page, limit)
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
	j := r.GetJson()
	m := api_model.CreateUser{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	u, err := model.GetUserByName(m.Username)
	if err == nil && u.Id != 0 {
		Fail(r, code.RESPONSE_ERROR, "用户已存在")
	}

	m.Password = utils.EncryptPassword(m.Password)
	user := model.GadminUser{
		UserName:   m.Username,
		Password:   m.Password,
		NickName:   m.Nickname,
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
	j := r.GetJson()
	m := api_model.UpdateUser{}
	_ = j.ToStruct(&m)
	if e := gvalid.CheckStruct(m, nil); e != nil {
		Fail(r, code.ERROR_INVALID_PARAM, e.String())
		return
	}
	u, err := model.GetUserByName(m.Username)
	if u == nil || err != nil || u.Id == 0 {
		Fail(r, code.RESPONSE_ERROR, "用户不存在")
		return
	}
	umap := gdb.Map{}
	umap = j.ToMap()
	delete(umap, "password")

	if m.Password == "" {
		delete(umap, "password")
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(r, code.RESPONSE_ERROR, err.Error())
			return
		}
	} else {
		umap["password"] = utils.EncryptPassword(m.Password)
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
	data := r.GetJson()
	id := data.GetInt("id")
	if id < 1 {
		Fail(r, code.RESPONSE_ERROR)
		return
	}
	u := new(model.GadminUser)
	user, err := u.GetById(id)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
		return
	}
	if user.UserName == model.ADMIN_NAME {
		Fail(r, code.RESPONSE_ERROR, "无权限")
		return
	}
	res, _ := u.DeleteById(id)
	if res <= 0 {
		Fail(r, code.RESPONSE_ERROR)
		return
	}
	model.Enforcer.DeleteRolesForUser(user.UserName)
	Success(r, "success")
}
