package model

import "github.com/gogf/gf/g/database/gdb"

// GetAllRole 获取所有角色名称
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllRole() (gdb.Result, error) {
	return defDB.Table("role_name").All()
}

// GetRoleByRoleKey 根据角色唯一键获取角色
//
// createTime:2019年05月06日 15:53:08
// author:hailaz
func GetRoleByRoleKey(role string) (RoleConfig, error) {
	obj := RoleConfig{}
	err := defDB.Table("role_name").Where("role_key", role).Struct(&obj)
	return obj, err
}
