package api

import (
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/library/code"
)

type UserController struct {
	BaseController
}

func (c *UserController) Info() {
	u := c.GetUser()
	if u != nil {
		Success(c.Request, u.GetUserInfo())
	}
	Fail(c.Request, code.RESPONSE_ERROR, "获取用户信息失败")
}

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
func (c *UserController) Post() {
	data := c.Request.GetJson()
	username := data.GetString("user_name")
	nickname := data.GetString("nick_name")
	email := data.GetString("email")
	password := data.GetString("password")
	passwordconfirm := data.GetString("passwordconfirm")
	phone := data.GetString("phone")

	u, err := model.GetUserByName(username)
	if err != nil || u.Id != 0 {
		Fail(c.Request, code.RESPONSE_ERROR, "用户已存在")
	}
	if password == "" {
		Fail(c.Request, code.RESPONSE_ERROR, "密码为空")
	}
	if password != passwordconfirm {
		Fail(c.Request, code.RESPONSE_ERROR, "输入密码不一致")
	}
	addu := c.GetUser()
	var addUserId = 0
	if addu != nil {
		addUserId = addu.Id
	}
	user := model.GadminUser{UserName: username, Password: password, NickName: nickname, Email: email, Phone: phone, AddUserId: addUserId}
	uid, _ := user.Insert()
	if uid > 0 {
		Success(c.Request, "success")
	}

	glog.Debug(uid)
	glog.Debug(data.ToJsonString())
	Fail(c.Request, code.RESPONSE_ERROR)
}
func (c *UserController) Put() {
	data := c.Request.GetJson()
	username := data.GetString("user_name")
	nickname := data.GetString("nick_name")
	email := data.GetString("email")
	password := data.GetString("password")
	passwordconfirm := data.GetString("passwordconfirm")
	phone := data.GetString("phone")

	u, err := model.GetUserByName(username)
	if err != nil || u.Id == 0 {
		Fail(c.Request, code.RESPONSE_ERROR, "用户不存在")
	}
	umap := gdb.Map{}
	if nickname != u.NickName && nickname != "" {
		umap["nick_name"] = nickname
	}
	if email != u.Email && email != "" {
		umap["email"] = email
	}
	if phone != u.Phone && phone != "" {
		umap["phone"] = phone
	}
	if password == "" {
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	} else {
		if password != passwordconfirm {
			Fail(c.Request, code.RESPONSE_ERROR, "输入密码不一致")
		}
		umap["password"] = service.EncryptPassword(password)
		err := model.UpdateUserById(u.Id, umap)
		if err != nil {
			Fail(c.Request, code.RESPONSE_ERROR, err.Error())
		}
	}

	Success(c.Request, "success")
}
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
