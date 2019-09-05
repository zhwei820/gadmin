package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
)

func GetAllPolicyMap() map[string]model.GadminPolicyconfig {
	allPolicys, _ := GetPagedPolicyList(1, 99999)
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
func GetPagedPolicyList(page, page_size int) ([]model.GadminPolicyconfig, int) {
	defaultName := "未命名"
	if page < 1 {
		page = 1
	}
	policyList := make([]model.GadminPolicyconfig, 0)
	policys := model.Enforcer.GetPermissionsForUser("system")
	total := len(policys)
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
		policyList = append(policyList, p)
	}
	if page_size == -1 {
		return policyList, total
	}
	if len(policyList) < page*page_size {
		if len(policyList) > page_size {
			policyList = policyList[(page-1)*page_size:]
		}
	} else {
		policyList = policyList[(page-1)*page_size : (page-1)*page_size+page_size]
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

// ReSetPolicy 更新路由权限
//
// createTime:2019年04月29日 17:30:26
// author:hailaz
func ReSetPolicy(role string, rmap map[string]model.RolePolicy) {
	old := model.Enforcer.GetPermissionsForUser(role)
	for _, item := range old {
		glog.Debug(item)
		full := fmt.Sprintf("%v %v %v", item[0], item[1], item[2])
		if _, ok := rmap[full]; ok { //从待插入列表中删除已存在的路由
			delete(rmap, full)
			//} else { //删除不存在的旧路由
			//	model.Enforcer.DeletePermissionForUser(item[0], item[1], item[2])
			//	if role == "system" {
			//		p, _ := model.GetPolicyByFullPath(fmt.Sprintf("%v:%v", item[1], item[2]))
			//		if p.Id > 0 {
			//			_, _ = p.DeleteById(p.Id)
			//		}
			//	}
		}
	}
	for _, item := range rmap { //插入新路由
		model.Enforcer.AddPolicy(item.Role, item.Path, item.Act)
	}
}
