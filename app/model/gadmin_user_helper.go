package model

import (
	"errors"
	"github.com/gogf/gf/g/database/gdb"
)

type UserInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

// {
//     roles: ['admin'],
//     introduction: 'I am a super administrator',
//     avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//     name: 'Super Admin'
//   }
// GetUserInfo 获取用户信息
//
// createTime:2019年05月08日 16:53:24
// author:hailaz
func (u *GadminUser) GetUserInfo() UserInfo {
	info := UserInfo{}
	if u.UserName == ADMIN_NAME {
		info.Roles = []string{ADMIN_NAME}
	} else {
		info.Roles = Enforcer.GetRolesForUser(u.UserName)
	}

	info.Avatar = u.Avatar
	info.Introduction = u.Introduction
	info.Name = u.NickName

	return info
}

// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetUserByName(name string) (*GadminUser, error) {
	u := GadminUser{}
	err := defDB.Table("user").Where("user_name", name).Struct(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UpdateUserById 更新用户
//
// createTime:2019年05月08日 14:28:18
// author:hailaz
func UpdateUserById(id int, udmap gdb.Map) error {
	r, err := defDB.Table("user").Data(udmap).Where("id=?", id).Update()
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i < 0 {
		return errors.New("update fail")
	}
	return nil
}

// GetUserByPageLimt 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetUserByPageLimt(page, limit int) ([]GadminUser, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	total, _ := defDB.Table("user").Count()
	if total == 0 {
		return nil, 0
	}

	userList := make([]GadminUser, 0)
	if total < page*limit {
		if total < limit {
			page = 1
		}
	}
	r, err := defDB.Table("user").Limit((page-1)*limit, (page-1)*limit+limit).Select()
	if err != nil {
		return nil, 0
	}
	r.ToStructs(&userList)
	return userList, total

}
