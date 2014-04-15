package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Bbs struct {
	Id       int64
	Userid   int64     `orm:"index"`
	Author   string    `orm:"size(15)"`
	Title    string    `orm:"size(100)"`
	Color    string    `orm:"size(7)"`
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"auto_now_add;type(datetime);index"`
	Views    int64
	Comments int64
	Status   int8
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Istop    int8
}

func (m *Bbs) TableName() string {
	return TableName("Bbs")
}

func (m *Bbs) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Bbs) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Bbs) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Bbs) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Bbs) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//带颜色的标题
func (m *Bbs) ColorTitle() string {
	if m.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>", m.Color, m.Title)
	} else {
		return m.Title
	}
}
