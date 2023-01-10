package main

import (
	"strings"
)

type PrimaryKey int

type SecondaryKey string

type UserRow struct {
	ID   PrimaryKey   //ID must be unique
	Name SecondaryKey //Name can be non-unique
	Age  int
}

type IndexByID map[PrimaryKey]UserRow

type IndexByName map[SecondaryKey][]PrimaryKey

type UserDB struct {
	ByID   IndexByID
	ByName IndexByName
}

func NewUser() *UserDB {
	return &UserDB{
		ByID:   make(IndexByID),
		ByName: make(IndexByName),
	}
}

func (db *UserDB) Insert(name string, age int) {
	// TODO: answer here
	primaryKey := len(db.ByID) + 1
	db.ByID[PrimaryKey(primaryKey)] = UserRow{
		Name: SecondaryKey(name),
		Age:  age,
		ID:   PrimaryKey(primaryKey),
	}

	db.ByName[SecondaryKey(name)] = append((*db).ByName[SecondaryKey(name)], PrimaryKey(primaryKey))
}

func (db *UserDB) WhereByID(id PrimaryKey) *UserRow {
	m, ok := db.ByID[id]
	if !ok {
		return nil
	}
	return &m
}

func (db *UserDB) WhereByName(name SecondaryKey) []*UserRow {
	ids := db.ByName[name]
	rows := make([]*UserRow, len(ids))
	// TODO: answer here

	for i, val := range ids {
		rows[i] = db.WhereByID(val)
	}
	return rows
}

func (db *UserDB) WhereNameBeginsWith(name string) []*UserRow {
	rows := make([]*UserRow, 0)
	// TODO: answer here
	for key, val := range (*db).ByName {
		if strings.HasPrefix(string(key), name) {
			for _, pk := range val {
				rows = append(rows, db.WhereByID(pk))
			}
		}
	}
	return rows
}
