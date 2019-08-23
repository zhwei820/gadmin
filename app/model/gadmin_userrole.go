package model

import (
	"errors"
)

// GadminUserrole 表名：gadmin_userrole
// 由数据库自动生成的结构体
type GadminUserrole struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	RoleKey  string `json:"role_key" xorm:"not null index VARCHAR(255)"`
	UserName string `json:"user_name" xorm:"not null VARCHAR(255)"`
}

// TableName 获取表名
func (t *GadminUserrole) TableName() string {
	return "gadmin_userrole"
}

// Insert 插入一条记录
func (t *GadminUserrole) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_userrole", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminUserrole) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_userrole", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminUserrole) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_userrole", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminUserrole) GetById(id int) (GadminUserrole, error) {
	obj := GadminUserrole{}
	err := defDB.Table("gadmin_userrole").Where("id", id).Struct(&obj)
	return obj, err
}
