package models

import "time"

type Mutant struct {
	ID         int
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
}
