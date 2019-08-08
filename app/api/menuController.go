package api

import (
	"github.com/gogf/gf/g/database/gdb"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type MenuController struct {
	BaseController
}

// @Summary menu list
// @Description menu list
// @Tags auth
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Success 200 {string} string	"ok"
// @router /menu [get]
func (c *MenuController) Get() {
	page := c.Request.GetInt("page", 1)
	limit := c.Request.GetInt("limit", 10)
	var list struct {
		List  []model.MenuOut `json:"items"`
		Total int             `json:"total"`
	}
	list.List, list.Total = model.GetMenuList(page, limit)

	Success(c.Request, list)
}

//
// @Summary create menu
// @Description create menu
// @Tags user
// @Param   SignInInfo  body model.MenuOut true "MenuOut"
// @Success 200 {string} string	"ok"
// @router /menu [post]
func (c *MenuController) Post() {
	data := c.Request.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	model.InsertMenuWithMeta(gdb.List{
		{
			"name":        m.Name,
			"menu_path":   m.MenuPath,
			"component":   m.Component,
			"sort":        m.Sort,
			"parent_name": m.ParentName,
			"hidden":      m.Hidden,
			"redirect":    m.Redirect,
			"alwaysshow":  m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.GadminMenumeta.Title,
				"icon":    m.GadminMenumeta.Icon,
				"noCache": m.GadminMenumeta.Nocache,
			},
		},
	})
	Success(c.Request, "添加成功")
}

//
// @Summary update menu
// @Description update menu
// @Tags user
// @Param   SignInInfo  body model.MenuOut true "MenuOut"
// @Success 200 {string} string	"ok"
// @router /menu [put]
func (c *MenuController) Put() {
	data := c.Request.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	err := model.UpdateMenuByName(
		m.Name,
		gdb.Map{
			"menu_path":   m.MenuPath,
			"component":   m.Component,
			"sort":        m.Sort,
			"parent_name": m.ParentName,
			"hidden":      m.Hidden,
			"redirect":    m.Redirect,
			"alwaysshow":  m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.GadminMenumeta.Title,
				"icon":    m.GadminMenumeta.Icon,
				"noCache": m.GadminMenumeta.Nocache,
			},
		},
	)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	Success(c.Request, "修改成功")
}

// @Summary delete menu
// @Description delete menu
// @Tags auth
// @Param	name	query 	string	true		"name"
// @Success 200 {string} string	"ok"
// @router /menu [delete]
func (c *MenuController) Delete() {
	data := c.Request.GetJson()
	name := data.GetString("name")
	m, err := model.GetMenuByName(name)
	if err != nil {
		Fail(c.Request, code.RESPONSE_ERROR, err.Error())
	}
	if m.AutoCreate > 0 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	res, _ := m.DeleteById(m.Id)
	if res <= 0 {
		Fail(c.Request, code.RESPONSE_ERROR)
	}
	Success(c.Request, "Delete")
}
