package mongo

import (
	"auth/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDatabase(env *config.Config) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		env.MongoDB.User,
		env.MongoDB.Password,
		env.MongoDB.Host,
		env.MongoDB.Port)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.MongoDB.CtxExp)*time.Second)
	defer cancel()
	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	mongoErr := mongo.Ping(ctx, readpref.Primary())
	if mongoErr != nil {
		return nil, mongoErr
	}
	return mongo, nil

}
