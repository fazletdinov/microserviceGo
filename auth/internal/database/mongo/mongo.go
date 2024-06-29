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

var MongoClient *mongo.Client

func InitDatabase() error {
	env := config.ConfigEnvs
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		env.MongoDB.User,
		env.MongoDB.Password,
		env.MongoDB.Host,
		env.MongoDB.Port)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConfigEnvs.MongoDB.CtxExp)*time.Second)
	defer cancel()
	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	MongoClient = mongo
	mongoErr := MongoClient.Ping(ctx, readpref.Primary())
	if mongoErr != nil {
		return mongoErr
	}
	return nil

}
