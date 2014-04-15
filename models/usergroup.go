package models

import (
	"github.com/astaxie/beego/orm"
)

//用户组表模型
type Usergroup struct {
	Id         int64
	Groupname  string `orm:"size(100)"`
	Isadmin    int8
	Manageauth string `orm:"size(1000)"`
}

func (m *Usergroup) TableName() string {
	return TableName("usergroup")
}

func (m *Usergroup) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Usergroup) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Usergroup) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Usergroup) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Usergroup) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Usergroups() orm.QuerySeter {
	return orm.NewOrm().QueryTable("usergroup").OrderBy("-Id")
}
