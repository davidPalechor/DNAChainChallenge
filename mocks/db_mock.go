package mocks

import (
	"DNAChainChallenge/models"
	"DNAChainChallenge/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/mock"
)

type DBMock struct{}

type mockQuerySeter struct {
	mock.DoNothingQuerySetter
	name string
}

func NewDBMock() utils.DBOrmer {
	return &DBMock{}
}

func (D DBMock) Insert(obj interface{}) (int64, error) {
	return 0, nil
}

func (D DBMock) One(queryable orm.QuerySeter, obj interface{}, cols ...string) error {
	return nil
}

func (D DBMock) GetQueryTable(model string) orm.QuerySeter {
	return &mockQuerySeter{name: model}
}

func (D DBMock) Count(querytable orm.QuerySeter) (int64, error) {
	mockTable, _ := querytable.(*mockQuerySeter)
	if mockTable.name == models.MutantTableName {
		return 40, nil
	} else if mockTable.name == models.HumanTableName {
		return 100, nil
	}
	return 0, nil
}
