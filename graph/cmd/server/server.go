package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clauribeirodevjava/13-GraphQL/graph"
	"github.com/clauribeirodevjava/13-GraphQL/graph/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {
	username := "root"
	password := "root"
	host := "localhost:3306" // Endereço do servidor MySQL
	databaseName := "mysql"

	// Crie a string de conexão
	connectionString := username + ":" + password + "@tcp(" + host + ")/" + databaseName

	// Abra a conexão com o banco de dados
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error()) // Em vez de usar panic
	}
	defer db.Close()

	Category := database.NewCategory(db)
	Course := database.NewCourse(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryField: Category,
		CourseField:   Course,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil)) // Adicione a vírgula e nil como handler
}
