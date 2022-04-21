package utils

import "github.com/beego/beego/v2/client/orm"

type DBOrmer interface {
	Insert(obj interface{}) (int64, error)
	One(queryable orm.QuerySeter, obj interface{}, cols ...string) error
	GetQueryTable(model string) orm.QuerySeter
}

type DBWrapper struct {
	db orm.Ormer
}

func NewDBWrapper() DBOrmer {
	return &DBWrapper{db: orm.NewOrm()}
}

func (d *DBWrapper) GetQueryTable(model string) orm.QuerySeter {
	return d.db.QueryTable(model)
}

func (d *DBWrapper) Insert(obj interface{}) (int64, error) {
	return d.db.Insert(obj)
}

func (d *DBWrapper) One(querytable orm.QuerySeter, obj interface{}, cols ...string) error {
	return querytable.One(obj, cols...)
}
