package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service/service_model"
	"strings"
)

func GetAllRoleMap() map[string]service_model.GadminRolePolicy {
	allRoles, _ := GetPagedRoleList(1, 99999)
	res := make(map[string]service_model.GadminRolePolicy, 0)
	for _, item := range allRoles {
		res[item.RoleKey] = item
	}
	return res
}

func GetRoleNames(allRole map[string]service_model.GadminRolePolicy, RoleKeys []string) []string {
	res := make([]string, 0)
	for _, RoleKey := range RoleKeys {
		res = append(res, allRole[RoleKey].Name)
	}
	return res
}

// GetRoleList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedRoleList(page, pageSize int) ([]service_model.GadminRolePolicy, int) {
	defaultname := "未命名"

	roleList := make([]service_model.GadminRolePolicy, 0)
	roles := model.Enforcer.GetAllRoles()
	total := len(roles)
	r, _ := model.GetAllRole()
	pn := make([]model.GadminRoleconfig, 0)
	_ = r.ToStructs(&pn)

	allPolicyMap := GetAllPolicyMap()
	for _, item := range roles {

		p := service_model.GadminRolePolicy{RoleKey: item, Name: defaultname}
		for _, itempn := range pn {
			if itempn.RoleKey == item {
				p.Name = itempn.Name
				p.Descrption = itempn.Descrption
				p.Id = itempn.Id
				break
			}
		}
		policyList := model.Enforcer.GetPermissionsForUser(item)
		for _, item := range policyList {
			fullPath := fmt.Sprintf("%v:%v", item[1], item[2])
			p.PolicyKeys = append(p.PolicyKeys, fullPath)
		}
		p.PolicyNames = GetPolicyNames(allPolicyMap, p.PolicyKeys)
		roleList = append(roleList, p)
	}
	if pageSize == -1 {
		return roleList, total
	}
	if len(roleList) < page*pageSize {
		if len(roleList) > pageSize {
			roleList = roleList[(page-1)*pageSize:]
		}
	} else {
		roleList = roleList[(page-1)*pageSize : (page-1)*pageSize+pageSize]
	}
	return roleList, total
}

// GetRoleByUsername 根据用户名获取对应角色
//
// createTime:2019年05月08日 15:08:19
// author:hailaz
func GetRoleByUsername(userName string) []model.GadminRoleconfig {
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
	r, err := model.GetRoleByRoleKey(role)
	if err != nil || r.Id == 0 {
		return errors.New("not exist")
	}
	// 存在则更新
	r.Name = name
	i, err := r.Update()
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
	r, err := model.GetRoleByRoleKey(role)
	// 不存在插入新数据
	if err != nil || r.Id == 0 {
		res := model.Enforcer.AddGroupingPolicy("system", role)
		if !res {
			return errors.New("add to casbin fail")
		}
		r := model.GadminRoleconfig{}
		r.RoleKey = role
		r.Name = name
		id, _ := r.Insert()
		if id > 0 {
			return nil
		} else {
			return errors.New("add to db fail")
		}
	}
	return errors.New("already exist")
}

// DeleteRole 删除角色
//
// createTime:2019年05月07日 11:12:59
// author:hailaz
func DeleteRole(role string) error {
	p, err := model.GetRoleByRoleKey(role)
	if err != nil || p.Id == 0 {
		return errors.New("not exist")
	}
	model.Enforcer.DeleteRole(role)
	i, _ := p.DeleteById(p.Id)
	if i > 0 {
		return nil
	}
	return errors.New("delete fail")
}

// SetUserRole 设置用户角色
//
// createTime:2019年05月08日 15:22:05
// author:hailaz
func SetUserRole(userName string, roles []string) {
	model.Enforcer.DeleteRolesForUser(userName)
	for _, item := range roles {
		model.Enforcer.AddRoleForUser(userName, item)
	}
}

func SetPolicyByRole(Policys []string, RoleKey string) {
	var routerMap = make(map[string]model.RolePolicy)
	for _, item := range Policys {
		list := strings.Split(item, ":")
		path := list[0]
		act := list[1]
		routerMap[item] = model.RolePolicy{Role: RoleKey, Path: path, Act: act}
	}
	ReSetPolicy(RoleKey, routerMap)
}

func GetRoleByRolekey(roleKey string) (ret service_model.GadminRolePolicy, err error) {
	res, err := model.GetRoleByRoleKey(roleKey)
	if err != nil || res.Id == 0 {
		return ret, errors.New("not exist")
	}
	allPolicyMap := GetAllPolicyMap()
	ret = service_model.GadminRolePolicy{
		Id:      res.Id,
		RoleKey: res.RoleKey,
		Name:    res.Name,
	}
	policys := model.Enforcer.GetPermissionsForUser(res.RoleKey)
	for _, item := range policys {
		fullPath := fmt.Sprintf("%v:%v", item[1], item[2])
		ret.PolicyKeys = append(ret.PolicyKeys, fullPath)
	}
	ret.PolicyNames = GetPolicyNames(allPolicyMap, ret.PolicyKeys)
	return ret, nil
}
