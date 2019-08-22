package model

import (
	"errors"
)

// GadminRolepolicy 表名：gadmin_rolepolicy
// 由数据库自动生成的结构体
type GadminRolepolicy struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	RoleKey    string `json:"role_key" xorm:"not null index VARCHAR(255)"`
	PolicyPath string `json:"policy_path" xorm:"not null index VARCHAR(255)"`
}

// TableName 获取表名
func (t *GadminRolepolicy) TableName() string {
	return "gadmin_rolepolicy"
}

// Insert 插入一条记录
func (t *GadminRolepolicy) Insert() (int64, error) {
	r, err := defDB.Insert("gadmin_rolepolicy", t)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	t.Id = int(id)
	return id, err
}

// Update 更新对象
func (t *GadminRolepolicy) Update() (int64, error) {
	if t.Id <= 0 {
		return 0, errors.New("primary_key <= 0")
	}
	r, err := defDB.Update("gadmin_rolepolicy", t, "id=?", t.Id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// DeleteById 删除一条记录
func (t *GadminRolepolicy) DeleteById(id int) (int64, error) {
	r, err := defDB.Delete("gadmin_rolepolicy", "id=?", id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

// GetById 通过id查询记录
func (t *GadminRolepolicy) GetById(id int) (GadminRolepolicy, error) {
	obj := GadminRolepolicy{}
	err := defDB.Table("gadmin_rolepolicy").Where("id", id).Struct(&obj)
	return obj, err
}
