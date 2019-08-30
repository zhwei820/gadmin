package service

import (
	"errors"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"strings"
)

// GetPagedRoleList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedRoleList(page, limit int) ([]model.GadminRoleconfig, int) {

	total, _ := model.CountRoleConfig()
	r, _ := model.GetPagedRoleConfig((page-1)*limit, limit)
	pcs := make([]model.GadminRoleconfig, 0)
	r.ToStructs(&pcs)

	return pcs, total
}

// GetRoleKeysByUserName 根据用户名获取对应角色
//
// createTime:2019年05月08日 15:08:19
// author:hailaz
func GetRoleKeysByUserName(userName string) []string {
	user, _ := model.GetUserByName(userName)

	return strings.Split(user.RoleKeys, ",")
}

// UpdateRoleByRoleKey 更新角色信息
//
// createTime:2019年05月06日 15:47:35
// author:hailaz
func UpdateRoleByRoleKey(role, name string) error {
	p, err := model.GetRoleByRoleKey(role)
	// 不存在报错
	if err != nil {
		return errors.New("update fail")
	}
	// 存在则更新
	p.Name = name
	i, err := p.Update()
	if err != nil {
		glog.Error(err)
		return err
	}
	if i < 0 {
		return errors.New("update fail")
	}
	return nil
}

// AddRole 添加角色
//
// createTime:2019年05月07日 10:45:04
// author:hailaz
func AddRole(role, name string) error {
	p, err := model.GetRoleByRoleKey(role)
	// 不存在插入新数据
	if err != nil || p.Id == 0 {
		res := model.Enforcer.AddGroupingPolicy("system", role)
		if !res {
			return errors.New("add fail")
		}
		p.RoleKey = role
		p.Name = name
		id, _ := p.Insert()
		if id > 0 {
			return nil
		} else {
			return errors.New("add fail")
		}
	}
	return errors.New("add fail")
}

// DeleteRole 删除角色
//
// createTime:2019年05月07日 11:12:59
// author:hailaz
func DeleteRole(role string) error {
	item, err := model.GetRoleByRoleKey(role)
	if err != nil || item.Id == 0 {
		return errors.New("delete fail")
	}
	model.Enforcer.DeleteRole(role)
	model.DeleteRolePolicys(role)
	i, _ := item.DeleteById(item.Id)
	if i > 0 {
		return nil
	}
	return errors.New("delete fail")
}

// SetRoleByUserName 设置用户角色
//
// createTime:2019年05月08日 15:22:05
// author:hailaz
func SetRoleByUserName(userName string, roles []string) error {
	model.Enforcer.DeleteRolesForUser(userName)
	for _, role := range roles {
		model.Enforcer.AddRoleForUser(userName, role)
	}
	rolestr := strings.Join(roles, ",")
	user, err := model.GetUserByName(userName)
	if err != nil {
		return errors.New("userName find User error")
	}
	user.RoleKeys = rolestr
	_, _ = user.Update()
	return nil
}
