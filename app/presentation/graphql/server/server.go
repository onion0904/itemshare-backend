package server

import (
	"log"
	"net/http"

	"github.com/onion0904/CarShareSystem/app/config"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db/migrations"
	mail_Service "github.com/onion0904/CarShareSystem/app/infrastructure/mail"
	mymiddleware "github.com/onion0904/CarShareSystem/app/middleware"
	"github.com/onion0904/CarShareSystem/app/presentation/graphql/graph"
	usecase_mail "github.com/onion0904/CarShareSystem/app/usecase/mail"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func Start() {
	cfg := config.GetConfig()
	DB := db.NewMainDB(cfg.DB)
	migrations.Migrate(DB)

	Port := cfg.Server.Port
	log.Printf("%s",Port)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DB: DB,
			EmailUseCase: usecase_mail.NewSendEmailUseCase(mail_Service.NewMailRepository()),
		},
		Directives: graph.Directive,
	}))

	// CORS対応。
	srv.AddTransport(transport.Options{})
	// GET / POST：GraphQLクエリを HTTP 経由で受け取るため。
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	
	// GraphQLクエリのキャッシュを有効化。LRU（最近使っていないものから削除）
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	
	// GraphQLのスキーマをクエリで確認できる「introspection」機能をオンにしてる。GraphQL Playgroundなどでスキーマの構造が見えるのはこれのおかげ。
	srv.Use(extension.Introspection{})
	// クライアントが「ハッシュ化されたクエリ」を使って通信することをサポート（帯域の節約になる）。
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", mymiddleware.AuthMiddleware(srv))

	log.Printf("Listening on 0.0.0.0:%s", Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+Port, nil))
}