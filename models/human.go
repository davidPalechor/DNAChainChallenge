package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

const HumanTableName = "human"

type Human struct {
	DNA        string    `orm:"column(dna);pk"`
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (h *Human) TableName() string {
	return HumanTableName
}

func init() {
	orm.RegisterModel(new(Human))
}
