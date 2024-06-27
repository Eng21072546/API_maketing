package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Ctx context.Context
var Cancel context.CancelFunc
var Client *mongo.Client
var Err error

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	Ctx, Cancel = context.WithTimeout(context.Background(),
		120*time.Second)
	Client, Err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))

	return Client, Ctx, Cancel, Err
}
