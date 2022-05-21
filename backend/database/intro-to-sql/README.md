# How to run non checkpoint code

Run the following code in the terminal:
1. First run `0-migration` inside root directory `intro-to-sql`
```
cd .../intro-to-sql
go run 0-migration/main.go
```

2. To run another code, run `go run name-of-code/main.go name-of-code/sql-name.go` example:
```
go run 1-create-table/main.go 1-create-table/sql-create-table.go
or
go run 2-read-data/main.go 2-read-data/sql-read-data.go
```

nb : You must run inside the root directory of the project `intro-to-sql`.
