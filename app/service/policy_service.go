package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
)

// GetPagedPolicyList 获取权限列表
//
// createTime:2019年05月06日 17:24:12
// author:hailaz
func GetPagedPolicyList(page, limit int) ([]model.GadminPolicyconfig, int) {
	if page < 1 {
		page = 1
	}
	total, _ := model.CountPolicyConfig()
	r, _ := model.GetPagedPolicyConfig((page-1)*limit, limit)
	pcs := make([]model.GadminPolicyconfig, 0)
	r.ToStructs(&pcs)

	return pcs, total
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
					p.DeleteById(p.Id)
				}
			}
		}
	}
	for _, item := range rmap { //插入新路由
		model.Enforcer.AddPolicy(item.Role, item.Path, item.Atc)
	}
}
