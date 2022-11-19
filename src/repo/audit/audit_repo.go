package audit

import (
	"audit-backend/container"
	"audit-backend/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func Initialize(r *repo.Database) *Repository {
	res := new(Repository)
	res.Collection = r.Client.Database(container.AppConfig.DatabaseName).Collection("events")
	return res
}
