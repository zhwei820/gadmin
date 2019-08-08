package model

import (
	"errors"
)

// MenuMeta 表名：menu_meta
// 由数据库自动生成的结构体
type MenuMeta struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	MenuName string `json:"menu_name" xorm:"not null unique VARCHAR(255)"`
	Title    string `json:"title" xorm:"not null VARCHAR(255)"`
	Icon     string `json:"icon" xorm:"VARCHAR(255)"`
	Nocache  int    `json:"nocache" xorm:"INT(11)"`
}

// TableName 获取表名
func (t *MenuMeta) TableName() string {
	return "menu_meta"
}

// Insert 插入一条记录
func (t *MenuMeta) Insert() (int64, error) {
	r, err := defDB.Insert("menu_meta", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = id
	return id, err
}

// Update 更新对象
func (t *MenuMeta) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("menu_meta", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *MenuMeta) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("menu_meta", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *MenuMeta) GetById(id int64) (MenuMeta, error) {
	obj := MenuMeta{}
	err := defDB.Table("menu_meta").Where("id", id).Struct(&obj)
	return obj, err
}
