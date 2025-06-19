package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"movie_api/config"
	"movie_api/graph"
	"movie_api/handlers"
	"movie_api/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// Connect to DB
	config.ConnectDB()

	// Buat Gin router
	r := gin.Default()

	// Static file route (untuk akses gambar)
	r.Static("/uploads", "./uploads")

	// Upload endpoint
	r.POST("/upload/poster", handlers.UploadHandler("posters"))
	r.POST("/upload/actor", handlers.UploadHandler("actors"))

	// GraphQL server setup
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000) )
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
    Cache: lru.New[string](100) ,
})

	// Playground
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL endpoint with middleware JWT
	r.POST("/query", func(c *gin.Context) {
		middleware.MiddlewareJWT(srv).ServeHTTP(c.Writer, c.Request)
	})

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("🚀 GraphQL Playground running at http://localhost:%s/", port)
	r.Run(":" + port)
}