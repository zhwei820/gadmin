package model

import "github.com/gogf/gf/g/database/gdb"

type RolePolicy struct {
	Role string
	Path string
	Atc  string
}

// GetAllRole 获取所有角色名称
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllRole() (gdb.Result, error) {
	return defDB.Table("gadmin_roleconfig").All()
}

// GetRoleByRoleKey 根据角色唯一键获取角色
//
// createTime:2019年05月06日 15:53:08
// author:hailaz
func GetRoleByRoleKey(role string) (GadminRoleconfig, error) {
	obj := GadminRoleconfig{}
	err := defDB.Table("gadmin_roleconfig").Where("role_key", role).Struct(&obj)
	return obj, err
}

func CountRoleConfig() (int, error) {
	return defDB.Table("gadmin_roleconfig").Count()
}

// GetPagedRoleConfig 获取分页的角色
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetPagedRoleConfig(limit ...int) (gdb.Result, error) {
	return defDB.Table("gadmin_roleconfig").Limit(limit...).Select()
}
