package audit

import (
	"audit-backend/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Repository struct {
	Collection *mongo.Collection
}

func Initialize(r *repo.Database) *Repository {
	res := new(Repository)
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		return nil
	}
	res.Collection = r.Client.Database(dbName).Collection("events")
	return res
}
