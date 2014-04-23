package models

import (
	"github.com/astaxie/beego/orm"
)

//用户组表模型
type Node struct {
	Id       int64
	Nodename string `orm:"size(100)"`
	Ismust   int32  //是否必选节点
}

func (m *Node) TableName() string {
	return TableName("node")
}

func (m *Node) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Node) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Node) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Node) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Node) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Nodes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("node").OrderBy("-Id")
}
