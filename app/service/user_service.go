package service

import (
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service/service_model"
)

// GetPagedUser 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetPagedUser(where map[string]interface{}, roleKey string, page, pageSize int) ([]service_model.GadminUserOut, int) {
	if roleKey != "" {
		usernames := model.Enforcer.GetUsersForRole(roleKey)
		where["username in (?)"] = usernames
	}
	userList := make([]service_model.GadminUserOut, 0)
	total, r, err := model.GetPagedUser(where, (page-1)*pageSize, pageSize)
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
