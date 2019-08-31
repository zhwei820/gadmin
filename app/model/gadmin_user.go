package model

import (
	"time"

	"errors"
)

// GadminUser 表名：gadmin_user
// 由数据库自动生成的结构体
type GadminUser struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Status       int       `json:"status" xorm:"not null INT(11)"`
	UserName     string    `json:"user_name" xorm:"not null unique VARCHAR(255)"`
	NickName     string    `json:"nick_name" xorm:"not null VARCHAR(255)"`
	Password     string    `json:"password" xorm:"not null VARCHAR(255)"`
	Email        string    `json:"email" xorm:"VARCHAR(255)"`
	Phone        string    `json:"phone" xorm:"VARCHAR(255)"`
	Sex          int       `json:"sex" xorm:"not null INT(11)"`
	CreateTime   time.Time `json:"create_time" xorm:"not null DATETIME(6)"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null DATETIME(6)"`
	AddUserId    int       `json:"add_user_id" xorm:"not null INT(11)"`
	Introduction string    `json:"introduction" xorm:"VARCHAR(255)"`
	Avatar       string    `json:"avatar" xorm:"VARCHAR(255)"`
}

// TableName 获取表名
func (t *GadminUser) TableName() string {
	return "gadmin_user"
}

// Insert 插入一条记录
func (t *GadminUser) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_user", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminUser) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_user", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminUser) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_user", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminUser) GetById(id int) (GadminUser, error) {
	obj := GadminUser{}
	err := defDB.Table("gadmin_user").Where("id", id).Struct(&obj)
	return obj, err
}
