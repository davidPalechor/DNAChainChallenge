package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

const MutantTableName = "mutant"

type Mutant struct {
	DNA        string    `orm:"column(dna);pk"`
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Mutant) TableName() string {
	return MutantTableName
}

func init() {
	orm.RegisterModel(new(Mutant))
}
