package logic

import (
	"DNAChainChallenge/models"
	"DNAChainChallenge/utils"
	"DNAChainChallenge/viewmodels"
)

type StatsLogic struct {
	DB utils.DBOrmer
}

func NewStatsLogic(db utils.DBOrmer) *StatsLogic {
	return &StatsLogic{
		DB: db,
	}
}

func (s *StatsLogic) Stats() (viewmodels.StatsResponse, error) {
	mutantTable := s.DB.GetQueryTable(models.MutantTableName)
	mutantCount, err := s.DB.Count(mutantTable)

	if err != nil {
		return viewmodels.StatsResponse{}, err
	}

	humanTable := s.DB.GetQueryTable(models.HumanTableName)
	humanCount, err := s.DB.Count(humanTable)

	if err != nil {
		return viewmodels.StatsResponse{}, err
	}

	response := viewmodels.StatsResponse{CountMutantDNA: mutantCount,
		CountHumanDNA: humanCount,
		Ratio:         float32(mutantCount) / float32(humanCount)}
	return response, err
}
