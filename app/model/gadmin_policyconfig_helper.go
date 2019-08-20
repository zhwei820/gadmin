package model

import (
	"github.com/gogf/gf/g/database/gdb"
)

// GetAllPolicyConfig 获取所有权限名称
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllPolicyConfig() (gdb.Result, error) {
	return defDB.Table("gadmin_policyconfig").All()
}

// GetPolicyByFullPath 根据权限全路径获取权限
//
// createTime:2019年05月06日 15:53:08
// author:hailaz
func GetPolicyByFullPath(path string) (GadminPolicyconfig, error) {
	obj := GadminPolicyconfig{}
	err := defDB.Table("gadmin_policyconfig").Where("full_path", path).Struct(&obj)
	return obj, err
}

func CountPolicyConfig() (int, error) {
	return defDB.Table("gadmin_policyconfig").Count()
}
