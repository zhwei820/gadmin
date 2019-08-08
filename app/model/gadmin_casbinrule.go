package model

import (
	"errors"
)

// GadminCasbinrule 表名：gadmin_casbinrule
// 由数据库自动生成的结构体
type GadminCasbinrule struct {
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
func (t *GadminCasbinrule) TableName() string {
	return "gadmin_casbinrule"
}

// Insert 插入一条记录
func (t *GadminCasbinrule) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_casbinrule", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminCasbinrule) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_casbinrule", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminCasbinrule) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_casbinrule", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminCasbinrule) GetById(id int) (GadminCasbinrule, error) {
	obj := GadminCasbinrule{}
	err := defDB.Table("gadmin_casbinrule").Where("id", id).Struct(&obj)
	return obj, err
}
