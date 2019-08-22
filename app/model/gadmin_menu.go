package model

import (
	"errors"
)

// GadminMenu 表名：gadmin_menu
// 由数据库自动生成的结构体
type GadminMenu struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	MenuPath   string `json:"menu_path" xorm:"not null VARCHAR(255)"`
	Component  string `json:"component" xorm:"not null VARCHAR(255)"`
	Redirect   string `json:"redirect" xorm:"not null VARCHAR(255)"`
	Hidden     int    `json:"hidden" xorm:"INT(11)"`
	Alwaysshow int    `json:"alwaysshow" xorm:"INT(11)"`
	Sort       int    `json:"sort" xorm:"INT(11)"`
	AutoCreate int    `json:"auto_create" xorm:"INT(11)"`
	ParentId   int    `json:"parent_id" xorm:"not null index INT(11)"`
}

// TableName 获取表名
func (t *GadminMenu) TableName() string {
	return "gadmin_menu"
}

// Insert 插入一条记录
func (t *GadminMenu) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_menu", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminMenu) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_menu", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminMenu) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_menu", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminMenu) GetById(id int) (GadminMenu, error) {
	obj := GadminMenu{}
	err := defDB.Table("gadmin_menu").Where("id", id).Struct(&obj)
	return obj, err
}
