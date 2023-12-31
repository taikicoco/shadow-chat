package main

import (
	"log"
	"net/http"
	"os"
	"server/graphql/generated"
	"server/graphql/generated/model"
	"server/graphql/resolver"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "1323"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("Defaulting to port %s", defaultPort)
		port = defaultPort
	}

	resolver := &resolver.Resolver{
		MessageID: make(map[int64][]chan<- *model.Message),
		Mutex:     sync.Mutex{},
	}
	gc := generated.Config{Resolvers: resolver}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gc))
	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
