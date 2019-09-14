package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"strings"
)

func GetAllPolicyMap() map[string]model.GadminPolicyconfig {
	allPolicys, _ := GetPagedPolicyList("", 1, 99999)
	res := make(map[string]model.GadminPolicyconfig, 0)
	for _, item := range allPolicys {
		res[item.FullPath] = item
	}
	return res
}

func GetPolicyNames(allPolicy map[string]model.GadminPolicyconfig, fullPath []string) []string {
	res := make([]string, 0)
	for _, path := range fullPath {
		res = append(res, allPolicy[path].Name)
	}
	return res
}

// GetPolicyList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedPolicyList(search string, page, pageSize int) ([]model.GadminPolicyconfig, int) {
	defaultName := "未命名"
	if page < 1 {
		page = 1
	}
	policyList := make([]model.GadminPolicyconfig, 0)
	policys := model.Enforcer.GetPermissionsForUser("system")
	r, _ := model.GetAllPolicyConfig()
	pcs := make([]model.GadminPolicyconfig, 0)
	_ = r.ToStructs(&pcs)

	for _, item := range policys {
		full := fmt.Sprintf("%v:%v", item[1], item[2])
		p := model.GadminPolicyconfig{FullPath: full, Name: defaultName}
		for _, itempc := range pcs {
			if itempc.FullPath == full {
				p.Name = itempc.Name
				p.Descrption = itempc.Descrption
				p.Label = itempc.Label
				p.Id = itempc.Id
				break
			}
		}
		if search != "" && !(strings.Contains(p.FullPath, search) ||
			strings.Contains(p.Name, search) ||
			strings.Contains(p.Label, search)) {
			continue
		}

		policyList = append(policyList, p)
	}
	total := len(policyList)

	if pageSize == -1 {
		return policyList, total
	}
	if len(policyList) < page*pageSize {
		if len(policyList) > pageSize {
			policyList = policyList[(page-1)*pageSize:]
		}
	} else {
		policyList = policyList[(page-1)*pageSize : (page-1)*pageSize+pageSize]
	}
	return policyList, total
}

// GetPolicyByRole 根据角色获取权限
//
// createTime:2019年05月07日 11:35:33
// author:hailaz
func GetPolicyByRole(role string) []model.GadminPolicyconfig {
	policyList := make([]model.GadminPolicyconfig, 0)
	policys := model.Enforcer.GetPermissionsForUser(role)
	glog.Debug(policys)
	for _, item := range policys {
		full := fmt.Sprintf("%v:%v", item[1], item[2])
		p := model.GadminPolicyconfig{FullPath: full}
		policyList = append(policyList, p)
	}
	return policyList
}

// UpdatePolicyByFullPath 更新权限信息
//
// createTime:2019年05月06日 15:47:35
// author:hailaz
func UpdatePolicyByFullPath(path, name, label string) error {
	p, err := model.GetPolicyByFullPath(path)
	if err != nil || p.Id == 0 {
		return errors.New("not exist")
	}
	// 存在则更新
	p.Name = name
	p.Label = label
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

// DeletePolicy 删除策略
//
// createTime:2019年05月07日 11:12:59
// author:hailaz
func DeletePolicy(path string) error {
	p, err := model.GetPolicyByFullPath(path)
	if err != nil || p.Id == 0 {
		return errors.New("not exist")
	}
	model.Enforcer.DeletePermission(path)
	i, _ := p.DeleteById(p.Id)
	if i > 0 {
		return nil
	}
	return errors.New("delete fail")
}

// ReSetPolicy 更新权限
//
// createTime:2019年04月29日 17:30:26
// author:hailaz
func ReSetPolicy(role string, rmap map[string]model.RolePolicy) {
	old := model.Enforcer.GetPermissionsForUser(role)

	if role != "system" {

		for _, key := range old {
			if _, ok := rmap[fmt.Sprintf("%v:%v", key[1], key[2])]; !ok {
				model.Enforcer.DeletePermissionForUser(role, key[1], key[2])
			}
		}
	}

	for _, item := range rmap {
		model.Enforcer.AddPolicy(item.Role, item.Path, item.Act)
	}
}
