package model

import (
	"errors"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/gtime"
)

// createTime:2019年04月23日 17:14:22
// author:hailaz
func GetUserByName(name string) (*User, error) {
	u := User{}
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
func UpdateUserById(id int64, udmap gdb.Map) error {
	udmap["update_time"] = gtime.Now().String()
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
