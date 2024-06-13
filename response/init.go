package response

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context
var cancel context.CancelFunc
var client *mongo.Client
var err error

func Init() {
	fmt.Printf("Init product controller varable")
	ctx = configs.Ctx
	cancel = configs.Cancel
	client = configs.Client
	err = configs.Err
	if err != nil {
		fmt.Println(err)
	}
	//defer cancel()
	//defer client.Disconnect(ctx)
}
