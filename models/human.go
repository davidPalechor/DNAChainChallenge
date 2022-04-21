package models

import "time"

type Human struct {
	dna        string
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
