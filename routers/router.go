package routers

import (
	"DNAChainChallenge/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/mutant",
				beego.NSInclude(&controllers.MutantController{})))

	beego.AddNamespace(ns)
}
