package service

import "github.com/hailaz/gadmin/app/model"

// GetPagedUser 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetPagedUser(where map[string]interface{}, page, limit int) ([]model.GadminUser, int) {

	total, _ := model.CountUser()

	userList := make([]model.GadminUser, 0)

	r, err := model.GetPagedUser(where, (page-1)*limit, limit)
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&userList)
	return userList, total

}
