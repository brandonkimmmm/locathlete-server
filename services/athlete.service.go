package services

import (
	"context"
	"fmt"
	"locathlete-server/configs"
	"locathlete-server/ent"
	"log"
)

func FindPaginatedAtheletes(ctx context.Context) ([]*ent.Athlete, error) {
	athletes, err := configs.DBClient.Athlete.Query().Limit(50).All(ctx)

	if err != nil {
		fmt.Println("failed querying athletes: %w", err.Error())
		return nil, err
	}
	log.Println("returned athletes:", athletes)
	return athletes, err
}
