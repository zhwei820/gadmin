package model

import (
	"errors"
)

// CasbinRule 表名：casbin_rule
// 由数据库自动生成的结构体
type CasbinRule struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	PType string `json:"p_type" xorm:"not null VARCHAR(255)"`
	V0    string `json:"v0" xorm:"not null VARCHAR(255)"`
	V1    string `json:"v1" xorm:"not null VARCHAR(255)"`
	V2    string `json:"v2" xorm:"not null VARCHAR(255)"`
	V3    string `json:"v3" xorm:"not null VARCHAR(255)"`
	V4    string `json:"v4" xorm:"not null VARCHAR(255)"`
	V5    string `json:"v5" xorm:"not null VARCHAR(255)"`
}

// TableName 获取表名
func (t *CasbinRule) TableName() string {
	return "casbin_rule"
}

// Insert 插入一条记录
func (t *CasbinRule) Insert() (int64, error) {
	r, err := defDB.Insert("casbin_rule", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *CasbinRule) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("casbin_rule", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *CasbinRule) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("casbin_rule", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *CasbinRule) GetById(id int64) (CasbinRule, error) {
	obj := CasbinRule{}
	err := defDB.Table("casbin_rule").Where("id", id).Struct(&obj)
	return obj, err
}
