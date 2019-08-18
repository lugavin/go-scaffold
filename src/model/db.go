package model

import (
	"log"
	"time"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db       *sql.DB
	database *mongo.Database
)

func init() {
	//initMysql()
	initMongo()
}

func initMysql() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/reliablemq")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL!")
}

func initMongo() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	database = client.Database("reliablemq")

	log.Println("Connected to MongoDB!")
}
