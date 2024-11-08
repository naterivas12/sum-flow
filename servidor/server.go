package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
		_ "github.com/go-sql-driver/mysql"
	"github.com/kissmarkrivas/sumago/graph"
	"github.com/kissmarkrivas/sumago/graph/generated"
	"github.com/kissmarkrivas/sumago/graph/model"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:123456@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.LogMode(true)
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE suma_db")
	db.Exec("USE suma_db")
	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Suma{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
