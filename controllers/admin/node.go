package admin

import (
	"github.com/astaxie/beego/validation"
	"goblog/models"
	"strings"
)

type NodeController struct {
	baseController
}

//列表
func (this *NodeController) List() {
	var page int64
	var pagesize int64 = 10
	var list []*models.Node
	var node models.Node

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := node.Query().Count()
	if count > 0 {
		node.Query().OrderBy("ismust", "-id").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/admin/node/list", true).ToString()
	this.display()
}

//添加
func (this *NodeController) Add() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.isPost() {
		nodename := strings.TrimSpace(this.GetString("nodename"))
		ismust := int32(0)

		if this.GetString("ismust") == "0" {
			ismust = 0
		} else {
			ismust = 1
		}

		input["nodename"] = nodename

		valid := validation.Validation{}

		if v := valid.Required(nodename, "nodename"); !v.Ok {
			errmsg["nodename"] = "请输入节点名"
		}

		if len(errmsg) == 0 {
			var node models.Node
			node.Nodename = nodename
			node.Ismust = ismust

			if err := node.Insert(); err != nil {
				this.showmsg(err.Error())
			}
			this.Redirect("/admin/node/list", 302)
		}

	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.display()
}

//编辑
func (this *NodeController) Edit() {
	id, _ := this.GetInt("id")
	node := models.Node{Id: id}
	if err := node.Read(); err != nil {
		this.showmsg("节点不存在")
	}

	errmsg := make(map[string]string)

	if this.isPost() {
		nodename := strings.TrimSpace(this.GetString("nodename"))

		if this.GetString("ismust") == "0" {
			node.Ismust = 0
		} else {
			node.Ismust = 1
		}

		valid := validation.Validation{}

		if v := valid.Required(nodename, "nodename"); !v.Ok {
			errmsg["nodename"] = "请输入节点名"
		}

		node.Nodename = nodename

		if len(errmsg) == 0 {
			node.Update()
			this.Redirect("/admin/node/list", 302)
		}

	}
	this.Data["errmsg"] = errmsg
	this.Data["node"] = node
	this.display()
}

//删除
func (this *NodeController) Delete() {
	id, _ := this.GetInt("id")

	node := models.Node{Id: id}
	if node.Read() == nil {
		node.Delete()
	}

	this.Redirect("/admin/node/list", 302)
}
