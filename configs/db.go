package configs

import (
	"context"
	"fmt"
	//"github.com/Eng21072546/API_maketing/controller"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://thanaphatboo:Eng21072546@cluster0.0c1ujka.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

//// client instance
//var DB *mongo.Client = ConnectDB()
//
//func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
//	collection := client.Database("API market").Collection(collectionName)
//}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

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
		30*time.Second)
	Client, Err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	//controller.Init() //init parameter in controller
	return Client, Ctx, Cancel, Err
}

//// Get config parameter
//func GetCtx() context.Context {
//	return ctx
//}
//func GetCancle() context.CancelFunc {
//	return cancel
//}
//func GetClient() *mongo.Client {
//	return client
//}
//func GetErr() error {
//	return err
//}
