package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/onion0904/CarShareSystem/app/config"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

const maxRetries = 5
const delay = 5 * time.Second

var (
	once  sync.Once
	query *dbgen.Queries
	dbcon *sql.DB
)

// contextからQueriesを取得する。contextにQueriesが存在しない場合は、パッケージ変数からQueriesを取得する
func GetQuery(ctx context.Context) *dbgen.Queries {
	txq := getQueriesWithContext(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func SetQuery(q *dbgen.Queries) {
	query = q
}

func GetDB() *sql.DB {
	return dbcon
}

func SetDB(d *sql.DB) {
	dbcon = d
}

type CtxKey string

const (
	QueriesKey CtxKey = "queries"
)

func WithQueries(ctx context.Context, q *dbgen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, q)
}

func getQueriesWithContext(ctx context.Context) *dbgen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*dbgen.Queries)
	if !ok {
		return nil
	}
	return queries
}

func LocalNewMainDB(cnf config.DBConfig) *sql.DB {
	once.Do(func() {
		dbcon, err := localConnect(cnf.DB_HOST, cnf.DB_PORT, cnf.DB_USER, cnf.DB_PASSWORD, cnf.DB_NAME)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})

	return dbcon
}

// localDBに接続する：最大5回リトライする
func localConnect(host, port, user, password, dbname string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		// データベースに接続
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("could not connect to db: %v", err)
		log.Printf("retrying in %v seconds...", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("could not connect to db after %d attempts", maxRetries)
}
