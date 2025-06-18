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

func NewMainDB(cnf config.DBConfig) *sql.DB {
	once.Do(func() {
		dbcon, err := connect(cnf.DB_URL)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})

	return dbcon
}

// dbに接続する：最大5回リトライする
func connect(db_url string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", db_url)
		if err != nil {
			log.Fatal("Error connecting to the database:", err)
		}
		if err != nil {
			log.Printf("Error opening DB connection: %v", err)

			return nil, fmt.Errorf("could not open db: %w", err)
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
