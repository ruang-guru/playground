package db

type MemoryDB struct {
	tables map[TableName]Rows
}

func NewMemoryDB(tables map[TableName]Rows) *MemoryDB {
	return &MemoryDB{tables}
}

func (db *MemoryDB) Load(tableName TableName) (Rows, error) {
	return db.tables[tableName], nil
}

func (db *MemoryDB) Save(tableName TableName, rows Rows) error {
	db.tables[tableName] = rows
	return nil
}

func (db *MemoryDB) Delete(tableName TableName) error {
	return nil
}
