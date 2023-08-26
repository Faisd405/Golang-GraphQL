package main

import (
	"graphql-template/auth"
	"graphql-template/database"
	graph "graphql-template/graphql"
	"net/http"

	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize the database connection.
	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	// Setting up Gin
	r := gin.Default()

	r.Use(auth.Middleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
