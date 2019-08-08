package model

import (
	"github.com/gogf/gf/g/database/gdb"
	"time"

	"errors"
)

// User 表名：user
// 由数据库自动生成的结构体
type User struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Status       int       `json:"status" xorm:"not null INT(11)"`
	UserName     string    `json:"user_name" xorm:"not null unique VARCHAR(255)"`
	NickName     string    `json:"nick_name" xorm:"not null VARCHAR(255)"`
	Password     string    `json:"password" xorm:"not null VARCHAR(255)"`
	Email        string    `json:"email" xorm:"VARCHAR(255)"`
	Phone        string    `json:"phone" xorm:"VARCHAR(255)"`
	Sex          int       `json:"sex" xorm:"not null INT(11)"`
	Age          int       `json:"age" xorm:"not null INT(11)"`
	AddTime      time.Time `json:"add_time" xorm:"not null DATETIME(6)"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null DATETIME(6)"`
	AddUserId    int       `json:"add_user_id" xorm:"not null INT(11)"`
	Introduction string    `json:"Introduction" xorm:"VARCHAR(255)"`
	Avatar       string    `json:"avatar" xorm:"VARCHAR(255)"`
}

// TableName 获取表名
func (t *User) TableName() string {
	return "user"
}

// Insert 插入一条记录
func (t *User) Insert() (int64, error) {
	r, err := defDB.Insert("user", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *User) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("user", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *User) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("user", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *User) GetById(id int64) (User, error) {
	obj := User{}
	err := defDB.Table("user").Where("id", id).Struct(&obj)
	return obj, err
}

type UserOut struct {
	Id           int64  `json:"id"`             //
	Status       int    `json:"status"`         //
	UserName     string `json:"user_name"`      //
	NickName     string `json:"nick_name"`      //
	Email        string `json:"email"`          //
	Phone        string `json:"phone"`          //
	Sex          int    `json:"sex"`            //
	Age          int    `json:"age"`            //
	AddTime      string `json:"add_time"`       //
	UpdateTime   string `json:"update_time"`    //
	AddUserId    int64  `json:"add_user_id"`    //
	ThirdPartyId int64  `json:"third_party_id"` //
	Introduction string `json:"Introduction"`   //
	Avatar       string `json:"avatar"`         //
}

type UserInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

// GetUserByPageLimt 获取用户
//
// createTime:2019年05月07日 16:11:41
// author:hailaz
func GetUserByPageLimt(page, limit int) ([]UserOut, int) {
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

	userList := make([]UserOut, 0)
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

// GetAllUser 获取所有用户
//
// createTime:2019年04月30日 10:20:50
// author:hailaz
func GetAllUser() (gdb.Result, error) {
	return defDB.Table("user").All()
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
func (u *User) GetUserInfo() UserInfo {
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
