package model

import (
	"errors"
)

// GadminRoleconfig 表名：gadmin_roleconfig
// 由数据库自动生成的结构体
type GadminRoleconfig struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	RoleKey    string `json:"role_key" xorm:"not null unique VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null VARCHAR(255)"`
	Descrption string `json:"descrption" xorm:"VARCHAR(255)"`
}

// TableName 获取表名
func (t *GadminRoleconfig) TableName() string {
	return "gadmin_roleconfig"
}

// Insert 插入一条记录
func (t *GadminRoleconfig) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_roleconfig", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminRoleconfig) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_roleconfig", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminRoleconfig) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_roleconfig", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminRoleconfig) GetById(id int) (GadminRoleconfig, error) {
	obj := GadminRoleconfig{}
	err := defDB.Table("gadmin_roleconfig").Where("id", id).Struct(&obj)
	return obj, err
}
