package api

import (
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/api/api_model"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/library/code"
)

type UserController struct {
	BaseController
}

// @Summary user info
// @Description user info
// @Tags user
// @Success 200 {string} string	"ok"
// @router /user/info [get]
func (c *UserController) Info() {
	u := c.GetUser()
	if u != nil {
		Success(c.Request, u.GetUserInfo())
	}
	Fail(c.Request, code.RESPONSE_ERROR, "获取用户信息失败")
}

// @Summary user menu
// @Description user menu
// @Tags user
// @Param	RoleConfig	query 	string	false		"RoleConfig"
// @Success 200 {string} string	"ok"
// @router /user/menu [get]
func (c *UserController) Menu() {
	RoleConfig := c.Request.GetString("RoleConfig")
	if RoleConfig != "" {
		var list struct {
			Menus     []model.MenuOut `json:"menus"`
			RoleMenus []model.MenuOut `json:"role_menus"`
		}
		list.Menus = model.GetMenuByRoleConfig([]string{model.ADMIN_NAME})
		list.RoleMenus = model.GetMenuByRoleConfig([]string{RoleConfig})
		Success(c.Request, list)
	}
	u := c.GetUser()
	if u != nil {
		if u.UserName == model.ADMIN_NAME {
			Success(c.Request, model.GetMenuByRoleConfig([]string{model.ADMIN_NAME}))
		} else {
			Success(c.Request, model.GetMenuByRoleConfig(model.Enforcer.GetRolesForUser(u.UserName)))
		}

	}
	Fail(c.Request, code.RESPONSE_ERROR, "获取用户菜单失败")
}

// @Summary user list
// @Description user list
// @Tags user
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Success 200 {string} string	"ok"
// @router /user [get]
func (c *UserController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var userList struct {
		List  []model.GadminUser `json:"items"`
		Total int                `json:"total"`
	}
	userList.List, userList.Total = model.GetUserByPageLimt(page, limit)
	Success(c.Request, userList)
}

//
// @Summary CreateUser
// @Description CreateUser
// @Tags user
// @Param   CreateUser  body api_model.CreateUser true "CreateUser"
// @Success 200 {string} string	"ok"
// @router /user [post]
func (c *UserController) Post() {
	j := c.Request.GetJson()
	m := api_model.CreateUser{}
	j.ToStruct(&m)

	u, err := model.GetUserByName(m.Username)
	if err != nil || u.Id != 0 {
		Fail(c.Request, code.RESPONSE_ERROR, "用户已存在")
	}
	addu := c.GetUser()
	var addUserId = 0
	if addu != nil {
		addUserId = addu.Id
	}
	user := model.GadminUser{UserName: m.Username, Password: m.Password, NickName: m.Nickname, Email: m.Email, Phone: m.Phone, AddUserId: addUserId}
	uid, _ := user.Insert()
	if uid > 0 {
		Success(c.Request, "success")
	}

	glog.Debug(uid)
	glog.Debug(j.ToJsonString())
	Fail(c.Request, code.RESPONSE_ERROR)
}

//
// @Summary UpdateUser
// @Description UpdateUser
// @Tags user
// @Param   UpdateUser  body api_model.UpdateUser true "UpdateUser"
// @Success 200 {string} string	"ok"
// @router /user [put]
func (c *UserController) Put() {
	j := c.Request.GetJson()
	m := api_model.UpdateUser{}
	j.ToStruct(&m)

	u, err := model.GetUserByName(m.Username)
	if err != nil || u.Id == 0 {
		Fail(c.Request, code.RESPONSE_ERROR, "用户不存在")
	}
	umap := gdb.Map{}
	umap = j.ToMap()
	delete(umap, "password")

	if m.Password == "" {
		delete(umap, "password")
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	} else {
		if m.Password != m.Passwordconfirm {
			Fail(c.Request, code.RESPONSE_ERROR, "输入密码不一致")
		}
		umap["password"] = service.EncryptPassword(m.Password)
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}

	Success(c.Request, "success")
}

//
// @Summary delete user
// @Description delete user
// @Tags user
// @Param	id	query 	integer	true		"id"
// @Success 200 {string} string	"ok"
// @router /user [delete]
func (c *UserController) Delete() {
	data := c.Request.GetJson()
	id := data.GetInt("id")
	if id < 1 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	u := new(model.GadminUser)
	user, err := u.GetById(id)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	if user.UserName == model.ADMIN_NAME {
		Fail(c.Request, code.RESPONSE_ERROR, "无权限")
	}
	res, _ := u.DeleteById(id)
	if res <= 0 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	model.Enforcer.DeleteRolesForUser(user.UserName)
	Success(c.Request, "success")
}
