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
	initMenu()
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
		UserName: ADMIN_NAME,
		NickName: ADMIN_NICK_NAME,
		Password: utils.EncryptPassword(ADMIN_DEFAULT_PASSWORD),
	}
	ii, err := admin.Insert()
	glog.Debugfln("%v %v", ii, err)

}

// initMenu 初始化菜单数据
//
// createTime:2019年05月16日 15:39:54
// author:hailaz
func initMenu() {
	InsertMenuWithMeta(gdb.List{
		{
			"name":        "user",
			"menu_path":   "/user",
			"component":   "layout",
			"redirect":    "/user/list",
			"sort":        "0",
			"parent_name": "",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "user",
				"icon":    "user",
				"noCache": 0},
		},
		{
			"name":        "userList",
			"menu_path":   "list",
			"component":   "user/user",
			"sort":        "0",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "userList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "roleList",
			"menu_path":   "/role/list",
			"component":   "user/role",
			"sort":        "1",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "roleList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "policyList",
			"menu_path":   "/policy/list",
			"component":   "user/policy",
			"sort":        "2",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "policyList",
				"icon":    "",
				"noCache": 0},
		},
		{
			"name":        "menuList",
			"menu_path":   "/menu/list",
			"component":   "user/menu",
			"sort":        "3",
			"parent_name": "user",
			"auto_create": true,
			"meta": gdb.Map{
				"title":   "menuList",
				"icon":    "",
				"noCache": 0},
		},
	})

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
