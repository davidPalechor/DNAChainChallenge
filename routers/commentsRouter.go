package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["DNAChainChallenge/controllers:MutantController"] = append(beego.GlobalControllerRouter["DNAChainChallenge/controllers:MutantController"],
        beego.ControllerComments{
            Method: "IsMutant",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["DNAChainChallenge/controllers:StatsController"] = append(beego.GlobalControllerRouter["DNAChainChallenge/controllers:StatsController"],
        beego.ControllerComments{
            Method: "Stats",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
