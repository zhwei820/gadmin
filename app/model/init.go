package model

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/utils"
	"github.com/hailaz/gadmin/utils/userdefinedpolicy"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"strings"
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
	InitCasbin()
	initSystemPolicyConfig()
	initUserDefinedPolicyConfig()
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
	glog.Debugf("%v %v", ii, err)
}

// InitCasbin 初始化Casbin
//
// createTime:2019年04月23日 14:45:20
// author:hailaz
func InitCasbin() {
	a := NewAdapter(defDB)
	Enforcer = casbin.NewEnforcer("./config/rbac.conf", a)
	_ = Enforcer.LoadPolicy()
	Enforcer.AddPolicy(ADMIN_NAME, "*", ACTION_ALL)
	glog.Debug("InitCasbin end ")
}

// 初始化自定义权限
func initUserDefinedPolicyConfig() {
	policys := userdefinedpolicy.UserDefinedPolicy
	r, _ := GetAllPolicyConfig()
	pcs := make([]GadminPolicyconfig, 0)
	_ = r.ToStructs(&pcs)
	pcd := make(map[string]GadminPolicyconfig)
	for _, itempc := range pcs {
		pcd[itempc.FullPath] = itempc
	}
	for label, policyList := range policys {
		for _, item := range policyList {
			full := item[0] // fmt.Sprintf("%v:%v", item[1], item[2])
			_, ok := pcd[full]
			if ok {
				continue
			}
			p := GadminPolicyconfig{FullPath: full, Name: item[1], Label: label}
			_, _ = p.Insert()
			items := strings.Split(full, ":")
			if len(items) == 2 { // 自定义权限加入casbin
				Enforcer.AddPolicy("system", items[0], items[1])
			}
		}
	}
}

// 初始化系统权限(api接口权限)
func initSystemPolicyConfig() {
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

// 初始化角色
func initRoleConfig() {
	roles := Enforcer.GetAllRoles()
	r, _ := GetAllRole()
	pns := make([]GadminRoleconfig, 0)
	_ = r.ToStructs(&pns)

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
