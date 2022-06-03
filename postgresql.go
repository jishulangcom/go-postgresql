package postgresql


import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jishulangcom/go-config"
)

var DB *pgxpool.Pool

func NewDB(cnf *config.PostgreSqlCnfDto, poolCnf *config.PostgreSqlPoolCnfDto) {
	var err error

	if cnf == nil {
		cnf = &config.PostgreSqlCnf
	}

	if poolCnf == nil {
		poolCnf = &config.PostgreSqlPoolCnf
	}

	//
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ",
		cnf.Host, cnf.Port, cnf.User, cnf.Pwd, cnf.DbName,
	)

	//
	dsn += fmt.Sprintf(" pool_max_conns=%d",
		poolCnf.MaxConn,
	)

	DB, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
