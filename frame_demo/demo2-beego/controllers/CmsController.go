package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("Cms", c.Cms)
}

// @router /cms [get]
func (this *CMSController) Cms() {
	this.TplName = "cms.html"
}
