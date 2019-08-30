package model

import (
	"errors"
)

// GadminPolicyconfig 表名：gadmin_policyconfig
// 由数据库自动生成的结构体
type GadminPolicyconfig struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FullPath   string `json:"full_path" xorm:"not null unique VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null VARCHAR(255)"`
	Descrption string `json:"descrption" xorm:"default 'NULL' VARCHAR(255)"`
}

// TableName 获取表名
func (t *GadminPolicyconfig) TableName() string {
	return "gadmin_policyconfig"
}

// Insert 插入一条记录
func (t *GadminPolicyconfig) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_policyconfig", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminPolicyconfig) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_policyconfig", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminPolicyconfig) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_policyconfig", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminPolicyconfig) GetById(id int) (GadminPolicyconfig, error) {
	obj := GadminPolicyconfig{}
	err := defDB.Table("gadmin_policyconfig").Where("id", id).Struct(&obj)
	return obj, err
}
