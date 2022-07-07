package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeeDemo/controllers:CMSController"] = append(beego.GlobalControllerRouter["BeeDemo/controllers:CMSController"],
        beego.ControllerComments{
            Method: "Cms",
            Router: `/cms`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
