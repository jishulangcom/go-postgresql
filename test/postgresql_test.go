package test

import (
	"context"
	"fmt"
	"github.com/jishulangcom/go-config"
	"github.com/jishulangcom/go-postgresql"
	"testing"
)

func Test(t *testing.T) {
	cnf := config.PostgreSqlCnfDto{
		Host:    "127.0.0.1",
		Port:    5432,
		User:    "postgres",
		Pwd:     "example",
		DbName:  "postgres",
		Debug:   true,
		MaxConn: 10,
	}
	postgresql.NewDB(&cnf, nil)

	//
	queryrTest()

	defer postgresql.CloseDB()
}

type pgTablesDto struct {
	schemaname string
	tablename  string
	tableowner interface{}
	tablespace interface{}
}
func queryrTest()  {
	rows, err := postgresql.DB.Query(context.Background(), "select schemaname, tablename, tableowner, j from pg_tables where tablespace=$1", "pg_global")
	if err != nil {
		panic(err)
	}
	//延迟关闭rows
	defer rows.Close()

	for rows.Next() {
		row := pgTablesDto{}
		err := rows.Scan(&row.schemaname, &row.tablename, &row.tableowner, &row.tablespace)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Schemaname = %v, Tablename = %v, Tableowner = %v, Tablespace = %v\n",
			row.schemaname,
			row.tablename,
			row.tableowner,
			row.tablespace,
		)
	}
	fmt.Println(rows, err)
}
