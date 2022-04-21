package controllers

import (
	"DNAChainChallenge/logic"
	beego "github.com/beego/beego/v2/server/web"
)

// MutantController operations for Mutant
type MutantController struct {
	beego.Controller
}

// GetAll ...
// @Title GetAll
// @Description get Mutant
// @Success 200 {object} models.Mutant
// @Failure 403
// @router / [get]
func (c *MutantController) GetAll() {
	logic := logic.NewMutantLogic()
	dnaMap := []string{"acab", "babb", "cbab", "bbca"}
	response := logic.IsMutant(dnaMap)
	c.Data["json"] = map[string]bool{"result": response}
	c.ServeJSON()
}
