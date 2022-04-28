package controllers

import (
	"DNAChainChallenge/logic"
	"DNAChainChallenge/utils"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type StatsController struct {
	beego.Controller
}

// Stats returns the statistics between identified mutants and humans
// @Title Stats
// @Description get Stats
// @Success 200 {object} viewmodels.StatsResponse
// @Failure 404
// @router / [get]
func (c *StatsController) Stats() {
	defer c.ServeJSON()

	statsLogic := logic.NewStatsLogic(utils.NewDBWrapper())
	count, err := statsLogic.Stats()

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		return
	}

	c.Data["json"] = count
}
