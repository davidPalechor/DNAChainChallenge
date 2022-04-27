package tests

import (
	"DNAChainChallenge/logic"
	"DNAChainChallenge/mocks"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println(apppath, file)
	beego.TestBeegoInit(apppath)
}

func TestMutantLogic(t *testing.T) {
	db := mocks.NewDBMock()
	logicM := logic.NewMutantLogic(db)

	tests := []struct {
		name      string
		input     []string
		response_ bool
	}{
		{
			name:      "Should be mutant (4x4) No diagonal Vertical Chains",
			input:     []string{"ATGT", "AGTT", "ACGT", "ATTT"},
			response_: true,
		},
		{
			name:      "Should be mutant (4x4) No diagonal Horizontal Chains",
			input:     []string{"AAAA", "AGTT", "ACGT", "TTTT"},
			response_: true,
		},
		{
			name:      "Should not be mutant (4x4) diagonal",
			input:     []string{"ATGC", "CAGT", "TTAT", "AGAA"},
			response_: false,
		},
	}

	Convey("Test Mutant Logic", t, func() {
		for _, test := range tests {
			Convey(test.name, func() {
				So(logicM.IsMutant(test.input), ShouldEqual, test.response_)
			})
		}
	})
}
