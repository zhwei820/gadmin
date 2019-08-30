package service

import (
	"errors"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
)

// GetRoleList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedRoleList(page, limit int) ([]model.GadminRoleconfig, int) {
	defaultname := "未命名"
	if page < 1 {
		page = 1
	}
	roleList := make([]model.GadminRoleconfig, 0)
	roles := model.Enforcer.GetAllRoles()
	total := len(roles)
	r, _ := model.GetAllRole()
	pn := make([]model.GadminRoleconfig, 0)
	_ = r.ToStructs(&pn)

	for _, item := range roles {
		p := model.GadminRoleconfig{RoleKey: item, Name: defaultname}
		for _, itempn := range pn {
			if itempn.RoleKey == item {
				p.Name = itempn.Name
				p.Descrption = itempn.Descrption
				break
			}
		}
		roleList = append(roleList, p)
	}
	if limit == -1 {
		return roleList, total
	}
	if len(roleList) < page*limit {
		if len(roleList) < limit {
			roleList = roleList
		} else {
			roleList = roleList[(page-1)*limit:]
		}
	} else {
		roleList = roleList[(page-1)*limit : (page-1)*limit+limit]
	}
	return roleList, total
}

// GetRoleByUserName 根据用户名获取对应角色
//
// createTime:2019年05月08日 15:08:19
// author:hailaz
func GetRoleByUserName(userName string) []model.GadminRoleconfig {
	roles := model.Enforcer.GetRolesForUser(userName)
	roleList := make([]model.GadminRoleconfig, 0)
	for _, item := range roles {
		p := model.GadminRoleconfig{RoleKey: item}
		roleList = append(roleList, p)
	}
	return roleList
}

// UpdateRoleByRoleKey 更新角色信息
//
// createTime:2019年05月06日 15:47:35
// author:hailaz
func UpdateRoleByRoleKey(role, name string) error {
	p, err := model.GetRoleByRoleKey(role)
	// 不存在插入新数据
	if err != nil || p.Id == 0 {
		p.RoleKey = role
		p.Name = name
		id, _ := p.Insert()
		if id > 0 {
			return nil
		} else {
			return errors.New("update fail")
		}
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
	p, err := model.GetRoleByRoleKey(role)
	if err != nil || p.Id == 0 {
		return errors.New("delete fail")
	}
	model.Enforcer.DeleteRole(role)
	i, _ := p.DeleteById(p.Id)
	if i > 0 {
		return nil
	}
	return errors.New("delete fail")
}

// SetRoleByUserName 设置用户角色
//
// createTime:2019年05月08日 15:22:05
// author:hailaz
func SetRoleByUserName(userName string, roles []string) {
	model.Enforcer.DeleteRolesForUser(userName)
	for _, item := range roles {
		model.Enforcer.AddRoleForUser(userName, item)
	}
}
