package userdefinedpolicy

var UserDefinedPolicy = map[string][][]string{
	"系统管理": {{"sys_menu:sys_menu", "系统管理"}},
	"角色管理": {{"role_list:role_list", "角色列表"}, {"role_edit:role_edit", "角色编辑"}},
	"用户管理": {{"user_list:user_list", "用户列表"}, {"user_edit:user_edit", "用户编辑"}},
	"权限管理": {{"policy_list:policy_list", "权限列表"}, {"policy_edit:policy_edit", "权限编辑"}},
}
