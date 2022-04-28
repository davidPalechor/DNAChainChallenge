package viewmodels

type StatsResponse struct {
	CountMutantDNA int64   `json:"count_mutant_dna"`
	CountHumanDNA  int64   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}
