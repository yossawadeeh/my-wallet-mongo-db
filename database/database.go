package database

import (
	"log"
	"log/slog"
	"my-wallet-ntier-mongo/constant"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoDB *mongo.Database

func ConnectDB() error {
	slog.Info("Connecting database...")

	uri := os.Getenv(constant.MONGODB_URI)
	docs := "www.mongodb.com/docs/drivers/go/current/"

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("Cannot connect database ", "error", err)
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()dd

	MongoDB = client.Database(constant.DATABASE_NAME)
	slog.Info("Connect database successfully !")
	return err
}
