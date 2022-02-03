package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/hasura_auth/config"
	"github.com/hasura_auth/handler"
	"github.com/hasura_auth/log"
	"github.com/hasura_auth/middle_ware"
	"github.com/hasura_auth/resolver"
	"github.com/hasura_auth/schema"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config.InitConfig()
	log.InitLogger()
	logrus.Infof("init schame: \n %s", schema.GetRootSchema())
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})
	http.Handle("/query", middle_ware.LogMiddleWare(handler.Handle(graphqlSchema)))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	if err := http.ListenAndServe(config.GlobalConfig.Address, nil); err != nil {
		panic(err)
	}
}
