package model

import "github.com/gogf/gf/g/database/gdb"

// IsMenuExist 菜单是否存在
//
// createTime:2019年05月17日 11:04:45
// author:hailaz
func IsMenuExist(name string) bool {
	m := Menu{}
	defDB.Table("menu").Where("name", name).Struct(&m)
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
			defDB.Insert("menu", item)
			defDB.Insert("menu_meta", mate)
		}
	}
}
