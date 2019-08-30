package model

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/utils"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var defDB gdb.DB
var Enforcer *casbin.Enforcer

const (
	ACTION_GET             = "(GET)"
	ACTION_POST            = "(POST)"
	ACTION_PUT             = "(PUT)"
	ACTION_DELETE          = "(DELETE)"
	ACTION_ALL             = "(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)"
	ADMIN_NAME             = "admin" //超级管理员用户名
	ADMIN_NICK_NAME        = "超级管理员" //超级管理员显示名称
	ADMIN_DEFAULT_PASSWORD = "123"   //超级管理员默认密码
)

// InitModel 初始化数据模型
//
// createTime:2019年05月13日 09:47:08
// author:hailaz
func InitModel() {
	defDB = g.DB()
	defDB.SetDebug(true)
	initUser()
	initCasbin()
	initSystemAndUserDefinedPolicyConfig()
	initRoleConfig()
}

// initUser 初始化用户
//
// createTime:2019年04月23日 14:57:23
// author:hailaz
func initUser() {
	u, err := GetUserByName(ADMIN_NAME)
	if err == nil && u.Id != 0 {
		return
	}
	admin := GadminUser{
		UserName:   ADMIN_NAME,
		NickName:   ADMIN_NICK_NAME,
		Password:   utils.EncryptPassword(ADMIN_DEFAULT_PASSWORD),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	ii, err := admin.Insert()
	glog.Debugfln("%v %v", ii, err)

}

// initCasbin 初始化Casbin
//
// createTime:2019年04月23日 14:45:20
// author:hailaz
func initCasbin() {
	a := NewAdapter(defDB)
	Enforcer = casbin.NewEnforcer("./config/rbac.conf", a)
	Enforcer.LoadPolicy()
	//Enforcer.DeletePermissionsForUser("group_admin")
	Enforcer.AddPolicy(ADMIN_NAME, "*", ACTION_ALL)
	//Enforcer.AddGroupingPolicy("system", "user")

}

// 初始化系统权限(api接口权限) 和 用户自定义权限
func initSystemAndUserDefinedPolicyConfig() {
	policys := Enforcer.GetPermissionsForUser("system")
	r, _ := GetAllPolicyConfig()
	pcs := make([]GadminPolicyconfig, 0)
	_ = r.ToStructs(&pcs)
	pcd := make(map[string]GadminPolicyconfig)
	for _, itempc := range pcs {
		pcd[itempc.FullPath] = itempc
	}
	for _, item := range policys {
		full := fmt.Sprintf("%v:%v", item[1], item[2])
		_, ok := pcd[full]
		if ok {
			continue
		}
		p := GadminPolicyconfig{FullPath: full, Name: "未命名"}
		_, _ = p.Insert()
	}
}

func initRoleConfig() {
	roles := Enforcer.GetAllRoles()
	r, _ := GetAllRole()
	pns := make([]GadminRoleconfig, 0)
	r.ToStructs(&pns)

	pcd := make(map[string]GadminRoleconfig)
	for _, item := range pns {
		pcd[item.RoleKey] = item
	}
	for _, item := range roles {
		_, ok := pcd[item]
		if ok {
			continue
		}
		p := GadminRoleconfig{RoleKey: item, Name: "未命名"}
		_, _ = p.Insert()
	}
}
