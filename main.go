package main

import (
	"fmt"
	"net/http"
	// "database/sql"
	// "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
  )

// db, err := sql.Open("mysql", "user:password@/dbname")

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
})

func main() {

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
	  Schema: &Schema,
	  Pretty: true,
	})
	
	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	fmt.Println("Ready to serve")
  
	// and serve!
	http.ListenAndServe(":8080", nil)
  }