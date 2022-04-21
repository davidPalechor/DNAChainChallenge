package models

import "time"

type Mutant struct {
	dna        string
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
