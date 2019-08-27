package api

import (
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/code"
)

type MenuController struct {
	BaseController
}

// @Summary menu list
// @Description menu list
// @Tags menu
// @Param	page	query 	integer	false		"page"
// @Param	limit	query 	integer	false		"limit"
// @Success 200 {string} string	"ok"
// @router /rbac/menu [get]
func (c *MenuController) Get(r *ghttp.Request) {
	//r.Context().
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)
	var list struct {
		List  []model.MenuOut `json:"items"`
		Total int             `json:"total"`
	}
	list.List, list.Total = model.GetMenuList(page, limit)

	Success(r, list)
}

//
// @Summary create menu
// @Description create menu
// @Tags menu
// @Param   SignInInfo  body model.MenuOut true "MenuOut"
// @Success 200 {string} string	"ok"
// @router /rbac/menu [post]
func (c *MenuController) Post(r *ghttp.Request) {
	data := r.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	model.InsertMenuWithMeta(gdb.List{
		{
			"menu_path":  m.MenuPath,
			"component":  m.Component,
			"sort":       m.Sort,
			"parent_id":  m.ParentId,
			"hidden":     m.Hidden,
			"redirect":   m.Redirect,
			"alwaysshow": m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.GadminMenumeta.Title,
				"icon":    m.GadminMenumeta.Icon,
				"noCache": m.GadminMenumeta.Nocache,
			},
		},
	})
	Success(r, "添加成功")
}

//
// @Summary update menu
// @Description update menu
// @Tags menu
// @Param   SignInInfo  body model.MenuOut true "MenuOut"
// @Success 200 {string} string	"ok"
// @router /rbac/menu [put]
func (c *MenuController) Put(r *ghttp.Request) {
	data := r.GetJson()
	m := model.MenuOut{}
	data.ToStruct(&m)
	err := model.UpdateMenuById(
		m.GadminMenu.Id,
		gdb.Map{
			"menu_path":  m.MenuPath,
			"component":  m.Component,
			"sort":       m.Sort,
			"parent_id":  m.ParentId,
			"hidden":     m.Hidden,
			"redirect":   m.Redirect,
			"alwaysshow": m.Alwaysshow,
			"meta": gdb.Map{
				"title":   m.GadminMenumeta.Title,
				"icon":    m.GadminMenumeta.Icon,
				"noCache": m.GadminMenumeta.Nocache,
			},
		},
	)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
	}
	Success(r, "修改成功")
}

// @Summary delete menu
// @Description delete menu
// @Tags menu
// @Param	name	query 	string	true		"name"
// @Success 200 {string} string	"ok"
// @router /rbac/menu [delete]
func (c *MenuController) Delete(r *ghttp.Request) {
	name := r.GetString("name")
	m, err := model.GetMenuByName(name)
	if err != nil {
		Fail(r, code.RESPONSE_ERROR, err.Error())
	}
	if m.AutoCreate > 0 {
		Fail(r, code.RESPONSE_ERROR)
	}
	res, _ := m.DeleteById(m.Id)
	if res <= 0 {
		Fail(r, code.RESPONSE_ERROR)
	}
	Success(r, "Delete")
}
