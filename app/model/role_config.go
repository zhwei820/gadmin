package model

import (
	"errors"
)

// RoleConfig 表名：role_config
// 由数据库自动生成的结构体
type RoleConfig struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	RoleKey    string `json:"role_key" xorm:"not null unique VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null VARCHAR(255)"`
	Descrption string `json:"descrption" xorm:"VARCHAR(255)"`
}

// TableName 获取表名
func (t *RoleConfig) TableName() string {
	return "role_config"
}

// Insert 插入一条记录
func (t *RoleConfig) Insert() (int64, error) {
	r, err := defDB.Insert("role_config", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *RoleConfig) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("role_config", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *RoleConfig) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("role_config", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *RoleConfig) GetById(id int64) (RoleConfig, error) {
	obj := RoleConfig{}
	err := defDB.Table("role_config").Where("id", id).Struct(&obj)
	return obj, err
}
