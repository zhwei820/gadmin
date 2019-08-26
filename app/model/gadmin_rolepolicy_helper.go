package model

import "github.com/gogf/gf/g/database/gdb"

// DeleteRoleMenus description
//
// createTime:2019年05月21日 17:52:06
// author:hailaz
func DeleteRolePolicys(role string) {
	defDB.Delete("gadmin_rolepolicy", "role_key=?", role)
}

// SetRolePolicys description
//
// createTime:2019年05月21日 17:54:38
// author:hailaz
func SetRolePolicys(role string, policys []string) {
	DeleteRolePolicys(role)
	ms := make(gdb.List, 0)
	for _, item := range policys {
		ms = append(ms, gdb.Map{"role_key": role, "policy_name": item})
	}

	defDB.Table("gadmin_rolepolicy").Data(ms).Insert()
}
