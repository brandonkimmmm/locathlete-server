package configs

import (
	"context"
	"locathlete-server/ent"
	"locathlete-server/ent/migrate"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DBClient *ent.Client

func DbStart() {
	client, err := ent.Open("postgres", "host="+os.Getenv("DATABASE_HOST")+" port="+os.Getenv("DATABASE_PORT")+" user="+os.Getenv("DATABASE_USER")+" dbname="+os.Getenv("DATABASE_NAME")+" password="+os.Getenv("DATABASE_PASSWORD")+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	DBClient = client
	if err := DBClient.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("Failed creating shcema resources: %v", err)
	}
}

func DbClose() {
	DBClient.Close()
}
