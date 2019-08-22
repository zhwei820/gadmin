package model

import (
	"errors"
)

// GadminRolemenu 表名：gadmin_rolemenu
// 由数据库自动生成的结构体
type GadminRolemenu struct {
	Id     int `json:"id" xorm:"not null pk autoincr INT(11)"`
	MenuId int `json:"menu_id" xorm:"not null index INT(11)"`
	RoleId int `json:"role_id" xorm:"not null index INT(11)"`
}

// TableName 获取表名
func (t *GadminRolemenu) TableName() string {
	return "gadmin_rolemenu"
}

// Insert 插入一条记录
func (t *GadminRolemenu) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_rolemenu", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminRolemenu) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_rolemenu", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminRolemenu) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_rolemenu", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminRolemenu) GetById(id int) (GadminRolemenu, error) {
	obj := GadminRolemenu{}
	err := defDB.Table("gadmin_rolemenu").Where("id", id).Struct(&obj)
	return obj, err
}
