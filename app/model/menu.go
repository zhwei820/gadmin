package model

import (
	"errors"
)

// Menu 表名：menu
// 由数据库自动生成的结构体
type Menu struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	MenuPath   string `json:"menu_path" xorm:"not null VARCHAR(255)"`
	Component  string `json:"component" xorm:"not null VARCHAR(255)"`
	Redirect   string `json:"redirect" xorm:"not null VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null unique VARCHAR(255)"`
	Hidden     int    `json:"hidden" xorm:"INT(11)"`
	Alwaysshow int    `json:"alwaysshow" xorm:"INT(11)"`
	Sort       int    `json:"sort" xorm:"INT(11)"`
	ParentName string `json:"parent_name" xorm:"VARCHAR(255)"`
	AutoCreate int    `json:"auto_create" xorm:"INT(11)"`
}

// TableName 获取表名
func (t *Menu) TableName() string {
	return "menu"
}

// Insert 插入一条记录
func (t *Menu) Insert() (int64, error) {
	r, err := defDB.Insert("menu", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *Menu) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("menu", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *Menu) DeleteById(id int64) (int64, error) {
	r, err := defDB.Delete("menu", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *Menu) GetById(id int64) (Menu, error) {
	obj := Menu{}
	err := defDB.Table("menu").Where("id", id).Struct(&obj)
	return obj, err
}
