package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
)

// GetPolicyList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedPolicyList(page, limit int) ([]model.GadminPolicyconfig, int) {
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
				break
			}
		}
		policyList = append(policyList, p)
	}
	if limit == -1 {
		return policyList, total
	}
	if len(policyList) < page*limit {
		if len(policyList) < limit {
			policyList = policyList
		} else {
			policyList = policyList[(page-1)*limit:]
		}
	} else {
		policyList = policyList[(page-1)*limit : (page-1)*limit+limit]
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
func UpdatePolicyByFullPath(path, name string) error {
	p, err := model.GetPolicyByFullPath(path)
	// 不存在插入新数据
	if err != nil || p.Id == 0 {
		p := model.GadminPolicyconfig{}
		p.FullPath = path
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
		} else { //删除不存在的旧路由
			model.Enforcer.DeletePermissionForUser(item[0], item[1], item[2])
			if role == "system" {
				p, _ := model.GetPolicyByFullPath(fmt.Sprintf("%v:%v", item[1], item[2]))
				if p.Id > 0 {
					_, _ = p.DeleteById(p.Id)
				}
			}
		}
	}
	for _, item := range rmap { //插入新路由
		model.Enforcer.AddPolicy(item.Role, item.Path, item.Atc)
	}
}
