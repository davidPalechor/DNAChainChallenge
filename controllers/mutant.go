package controllers

import (
	"DNAChainChallenge/logic"
	"DNAChainChallenge/utils"
	"DNAChainChallenge/viewmodels"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

// MutantController operations for Mutant
type MutantController struct {
	beego.Controller
}

// IsMutant Returns whether a DNA chain corresponds to a mutant or human
// @Title IsMutant
// @Description post Mutant
// @Success 200
// @Failure 400
// @router / [post]
func (c *MutantController) IsMutant() {
	defer c.ServeJSON()

	var request viewmodels.IsMutantRequest

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = request
		return
	}

	logicM := logic.NewMutantLogic(utils.NewDBWrapper())
	isMutant := logicM.IsMutant(request.DNA)
	if !isMutant {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
	}

	if err := logicM.SaveChain(request, isMutant); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		return
	}
}
