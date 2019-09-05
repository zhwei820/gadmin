package model

import (
	"errors"
	"github.com/gogf/gf/g/database/gdb"
)

type UserRoles struct {
	Roles []string `json:"roles"`
}

// {
//     roles: ['admin'],
//     introduction: 'I am a super administrator',
//     avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//     name: 'Super Admin'
//   }
// GetUserRoles 获取用户信息
//
// createTime:2019年05月08日 16:53:24
// author:hailaz
func (u *GadminUser) GetUserRoles() UserRoles {
	info := UserRoles{}
	if u.Username == ADMIN_NAME {
		info.Roles = []string{ADMIN_NAME}
	} else {
		info.Roles = Enforcer.GetRolesForUser(u.Username)
	}

	return info
}

// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetUserByName(name string) (*GadminUser, error) {
	u := GadminUser{}
	err := defDB.Table("gadmin_user").Where("username", name).Struct(&u)
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
	r, err := defDB.Table("gadmin_user").Data(udmap).Where("id=?", id).Update()
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i < 0 {
		return errors.New("update fail")
	}
	return nil
}

func CountUser() (int, error) {
	return defDB.Table("gadmin_user").Count()
}

// GetPagedUser 获取分页的用户
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetPagedUser(where map[string]interface{}, page_size ...int) (gdb.Result, error) {
	qs := defDB.Table("gadmin_user")
	for key := range where {
		qs = qs.Where(key, where[key])
	}
	return qs.Limit(page_size...).Select()
}
