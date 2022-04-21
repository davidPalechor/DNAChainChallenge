package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Initial_20220421_105139 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Initial_20220421_105139{}
	m.Created = "20220421_105139"

	migration.Register("Initial_20220421_105139", m)
}

// Up Run the migrations
func (m *Initial_20220421_105139) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(
		`CREATE TABLE IF NOT EXISTS mutant (
				dna CHAR(64) NOT NULL PRIMARY KEY,
				inserted_at TIMESTAMP NOT NULL
			)`,
	)

	m.SQL(
		`CREATE TABLE IF NOT EXISTS human (
				dna CHAR(64) NOT NULL PRIMARY KEY,
				inserted_at TIMESTAMP NOT NULL
			)`,
	)
}

// Down Reverse the migrations
func (m *Initial_20220421_105139) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE mutant`)
	m.SQL(`DROP TABLE human`)
}
