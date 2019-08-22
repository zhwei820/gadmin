package model

import (
	"errors"
)

// GadminMenumeta 表名：gadmin_menumeta
// 由数据库自动生成的结构体
type GadminMenumeta struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title   string `json:"title" xorm:"not null VARCHAR(255)"`
	Icon    string `json:"icon" xorm:"VARCHAR(255)"`
	Nocache int    `json:"nocache" xorm:"INT(11)"`
	MenuId  int    `json:"menu_id" xorm:"not null index INT(11)"`
}

// TableName 获取表名
func (t *GadminMenumeta) TableName() string {
	return "gadmin_menumeta"
}

// Insert 插入一条记录
func (t *GadminMenumeta) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_menumeta", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminMenumeta) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_menumeta", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminMenumeta) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_menumeta", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminMenumeta) GetById(id int) (GadminMenumeta, error) {
	obj := GadminMenumeta{}
	err := defDB.Table("gadmin_menumeta").Where("id", id).Struct(&obj)
	return obj, err
}
