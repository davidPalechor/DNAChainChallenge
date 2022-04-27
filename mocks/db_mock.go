package mocks

import (
	"DNAChainChallenge/utils"
	"github.com/beego/beego/v2/client/orm"
)

type DBMock struct{}

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
	return nil
}
