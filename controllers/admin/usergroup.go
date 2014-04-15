package admin

import (
	"github.com/astaxie/beego/validation"
	"goblog/models"
	"strings"
)

type UsergroupController struct {
	baseController
}

//用户组列表
func (this *UsergroupController) List() {
	var page int64
	var pagesize int64 = 10
	var list []*models.Usergroup
	var usergroup models.Usergroup

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := usergroup.Query().Count()
	if count > 0 {
		usergroup.Query().OrderBy("isadmin", "-id").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/admin/usergroup/list", true).ToString()
	this.display()
}

//添加用户组
func (this *UsergroupController) Add() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.isPost() {
		groupname := strings.TrimSpace(this.GetString("groupname"))
		isadmin := int8(0)
		manageauth := strings.TrimSpace(this.GetString("manageauth"))

		if this.GetString("isadmin") == "0" {
			isadmin = 0
		} else {
			isadmin = 1
		}

		input["groupname"] = groupname
		input["manageauth"] = manageauth

		valid := validation.Validation{}

		if v := valid.Required(groupname, "groupname"); !v.Ok {
			errmsg["groupname"] = "请输入用户组名"
		}
		if v := valid.Required(manageauth, "manageauth"); !v.Ok {
			errmsg["manageauth"] = "请添加权限列表"
		}

		if len(errmsg) == 0 {
			var usergroup models.Usergroup
			usergroup.Groupname = groupname
			usergroup.Isadmin = isadmin
			usergroup.Manageauth = manageauth

			if err := usergroup.Insert(); err != nil {
				this.showmsg(err.Error())
			}
			this.Redirect("/admin/usergroup/list", 302)
		}

	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.display()
}

//编辑用户组
func (this *UsergroupController) Edit() {
	id, _ := this.GetInt("id")
	usergroup := models.Usergroup{Id: id}
	if err := usergroup.Read(); err != nil {
		this.showmsg("用户组不存在")
	}

	errmsg := make(map[string]string)

	if this.isPost() {
		groupname := strings.TrimSpace(this.GetString("groupname"))
		manageauth := strings.TrimSpace(this.GetString("manageauth"))

		if this.GetString("isadmin") == "0" {
			usergroup.Isadmin = 0
		} else {
			usergroup.Isadmin = 1
		}

		valid := validation.Validation{}

		if v := valid.Required(groupname, "groupname"); !v.Ok {
			errmsg["groupname"] = "请输入用户组名"
		}
		if v := valid.Required(manageauth, "manageauth"); !v.Ok {
			errmsg["manageauth"] = "请添加权限列表"
		}

		usergroup.Groupname = groupname
		usergroup.Manageauth = manageauth

		if len(errmsg) == 0 {
			usergroup.Update()
			this.Redirect("/admin/usergroup/list", 302)
		}

	}
	this.Data["errmsg"] = errmsg
	this.Data["usergroup"] = usergroup
	this.display()
}

//删除用户
func (this *UsergroupController) Delete() {
	id, _ := this.GetInt("id")

	usergroup := models.Usergroup{Id: id}
	if usergroup.Read() == nil {
		usergroup.Delete()
	}

	this.Redirect("/admin/usergroup/list", 302)
}
