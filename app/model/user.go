package model

import (
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
	t.Id = id
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
