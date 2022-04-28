package tests

import (
	"DNAChainChallenge/logic"
	"DNAChainChallenge/mocks"
	"DNAChainChallenge/viewmodels"
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

func TestStatsLogic(t *testing.T) {
	db := mocks.NewDBMock()
	logicS := logic.StatsLogic{DB: db}

	tests := []struct {
		name      string
		response_ viewmodels.StatsResponse
	}{
		{
			name:      "Should respond successfuly",
			response_: viewmodels.StatsResponse{CountHumanDNA: 100, CountMutantDNA: 40, Ratio: 0.4},
		},
	}

	Convey("Test Mutant Logic", t, func() {
		for _, test := range tests {
			Convey(test.name, func() {
				logicResponse, _ := logicS.Stats()
				So(logicResponse, ShouldResemble, test.response_)
			})
		}
	})
}
