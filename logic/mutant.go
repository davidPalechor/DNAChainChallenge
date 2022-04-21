package logic

import (
	"DNAChainChallenge/utils"
	"strings"
)

type MutanLogic struct {
	DB utils.DBOrmer
}

func NewMutantLogic() *MutanLogic {
	return &MutanLogic{}
}

func checkEqualItems(list [4]string) int {
	for i := 1; i < 4; i++ {
		if strings.ToLower(list[i]) != strings.ToLower(list[0]) {
			return 0
		}
	}
	return 1
}

func transpose(matrix [4][4]string) [][]string {
	transMatrix := make([][]string, 4)
	for i := range transMatrix {
		transMatrix[i] = make([]string, 4)
	}

	for i := 0; i < len(matrix); i++ {
		transRow := transMatrix[i]
		for j := 0; j < len(transRow); j++ {
			transRow[j] = matrix[j][i]
		}
	}
	return transMatrix
}

func checkDNAMap(matrix [4][4]string) int {
	var transRow [4]string
	count := 0
	transMatrix := transpose(matrix)
	for i := 0; i < 4; i++ {
		copy(transRow[:], transMatrix[i])
		count += checkEqualItems(matrix[i])
		count += checkEqualItems(transRow)
		if count == 2 {
			return count
		}
	}
	// Diagonals
	leftDiagonal := [4]string{
		matrix[0][0],
		matrix[1][1],
		matrix[2][2],
		matrix[3][3],
	}

	rightDiagonal := [4]string{
		matrix[0][3],
		matrix[1][2],
		matrix[2][1],
		matrix[3][0],
	}
	count += checkEqualItems(leftDiagonal)
	count += checkEqualItems(rightDiagonal)
	return count
}

func buildSubMatrix(matrix [][]string, row int, col int) [4][4]string {
	var subMatrix [4][4]string
	for i := 0; i < 4; i++ {
		copy(subMatrix[i][:], matrix[row][col:col+4])
		row++
	}
	return subMatrix
}

func buildMatrix(matrix []string) [][]string {
	resultMatrix := make([][]string, len(matrix))
	for j, codex := range matrix {
		resultMatrix[j] = strings.Split(codex, "")
	}
	return resultMatrix
}

func (m *MutanLogic) IsMutant(dna []string) bool {
	matrixLength := len(dna)
	count := 0

	if matrixLength < 4 {
		return false
	}

	dnaMatrix := buildMatrix(dna)
	if matrixLength == 4 {
		dnaSubMatrix := buildSubMatrix(dnaMatrix, 0, 0)
		if localCount := checkDNAMap(dnaSubMatrix); localCount >= 2 {
			return true
		}
		return false
	}

	for i := 0; i < matrixLength-3; i++ {
		for j := 0; j < matrixLength-3; j++ {
			dnaSubMatrix := buildSubMatrix(dnaMatrix, i, j)
			count += checkDNAMap(dnaSubMatrix)
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
