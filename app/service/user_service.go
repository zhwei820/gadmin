package service

import (
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service/service_model"
)

// GetPagedUser 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetPagedUser(where map[string]interface{}, page, page_size int) ([]service_model.GadminUserOut, int) {
	total, _ := model.CountUser()
	userList := make([]service_model.GadminUserOut, 0)

	r, err := model.GetPagedUser(where, (page-1)*page_size, page_size)
	if err != nil {
		return nil, 0
	}
	_ = r.ToStructs(&userList)

	allRoleMap := GetAllRoleMap()
	for ii := range userList {
		userList[ii].Password = ""
		userList[ii].Roles = GetRoleNames(allRoleMap, model.Enforcer.GetRolesForUser(userList[ii].Username))
	}
	return userList, total
}
