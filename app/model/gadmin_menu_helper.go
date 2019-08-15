package model

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/hailaz/gadmin/utils"
)

type MenuOut struct {
	GadminMenumeta
	GadminMenu
	Children []MenuOut `json:"children"`
}

// IsMenuExist 菜单是否存在
//
// createTime:2019年05月17日 11:04:45
// author:hailaz
func IsMenuExist(name string) bool {
	m := GadminMenu{}
	defDB.Table("gadmin_menu").Where("name", name).Struct(&m)
	if m.Id > 0 {
		return true
	}
	return false
}

// InsertMenuWithMeta 插入菜单
//
// createTime:2019年05月17日 11:12:13
// author:hailaz
func InsertMenuWithMeta(list gdb.List) {
	for _, item := range list {
		if !IsMenuExist(item["name"].(string)) {
			mate := item["meta"].(gdb.Map)
			mate["menu_name"] = item["name"].(string)
			delete(item, "meta")
			defDB.Insert("gadmin_menu", item)
			defDB.Insert("gadmin_menumeta", mate)
		}
	}
}

// GetMenuList 获取菜单列表
//
// createTime:2019年05月17日 16:17:33
// author:hailaz
func GetMenuList(page, limit int) ([]MenuOut, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	total, _ := defDB.Table("gadmin_menu").Count()
	if total == 0 {
		return nil, 0
	}

	menuList := make([]MenuOut, 0)
	if total < page*limit {
		if total < limit {
			page = 1
		}
	}
	r, err := defDB.Table("gadmin_menu").Limit((page-1)*limit, (page-1)*limit+limit).Select()
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&menuList)
	for index, item := range menuList {
		meta := GadminMenumeta{}
		r, _ := defDB.Table("gadmin_menumeta").Where("menu_name=?", item.Name).One()
		r.ToStruct(&meta)
		menuList[index].GadminMenumeta = meta
	}
	return menuList, total
}

// UpdateMenuByName description
//
// createTime:2019年05月17日 17:54:40
// author:hailaz
func UpdateMenuByName(name string, dataMap gdb.Map) error {
	mate := dataMap["meta"].(gdb.Map)
	delete(dataMap, "meta")
	_, err := defDB.Update("gadmin_menu", dataMap, "name=?", name)
	if err != nil {
		return err
	}
	_, err = defDB.Update("gadmin_menumeta", mate, "menu_name=?", name)
	if err != nil {
		return err
	}
	return nil
}

// GetMenuByRoleconfig description
//
// createTime:2019年05月16日 17:19:53
// author:hailaz
func GetMenuByRoleConfig(roles []string) []MenuOut {
	menus := make([]MenuOut, 0)
	if utils.IsStringInSlice(ADMIN_NAME, roles) {
		r, _ := defDB.Table("gadmin_menu").All()
		r.ToStructs(&menus)
	} else {
		roleSlice := make(g.Slice, 0)
		for _, item := range roles {
			roleSlice = append(roleSlice, item)
		}
		r, _ := defDB.Table("role_menu rm,menu m").Where("rm.menu_name=m.name AND rm.role_key IN (?)", roleSlice).Fields("m.*").All()
		r.ToStructs(&menus)
	}

	for index, item := range menus {
		meta := GadminMenumeta{}
		r, _ := defDB.Table("gadmin_menumeta").Where("menu_name=?", item.Name).One()
		r.ToStruct(&meta)
		menus[index].GadminMenumeta = meta
	}
	menuRoot := make([]MenuOut, 0)
	childs := make([]*MenuOut, 0)
	for index, item := range menus { //分类菜单，一级菜单与非一级菜单
		if item.ParentName == "" {
			menuRoot = append(menuRoot, item)
		} else {
			childs = append(childs, &menus[index])
		}
	}
	for index, _ := range menuRoot {
		FindChildren(&menuRoot[index], childs)
	}

	return menuRoot
}

// FindChildren 找子菜单
//
// createTime:2019年05月17日 09:15:52
// author:hailaz
func FindChildren(mo *MenuOut, list []*MenuOut) {
	for _, item := range list {
		if item.ParentName == mo.Name {
			mo.Children = append(mo.Children, *item)
		}
	}
	for index := 0; index < len(mo.Children); index++ {
		FindChildren(&mo.Children[index], list)
	}
}

// GetMenuByName 根据名称获取菜单
//
// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetMenuByName(name string) (*GadminMenu, error) {
	m := GadminMenu{}
	err := defDB.Table("gadmin_menu").Where("name", name).Struct(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
