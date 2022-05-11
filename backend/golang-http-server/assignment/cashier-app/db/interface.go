package db

type DBName = string
type TableName = string
type Rows = []Row
type Row = []Cell
type Cell = string

type DB interface {
	Load(dbName DBName) (Rows, error)
	Save(dbName DBName, rows Rows) error
	Delete(dbName DBName) error
}
