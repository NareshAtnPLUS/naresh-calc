package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MiniKartV1/calc/internal/models"
	"github.com/MiniKartV1/calc/internal/types"
	user_models "github.com/MiniKartV1/minikart-auth/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Adapter struct {
	db                *mongo.Database
	usersCollection   *mongo.Collection
	computeCollection *mongo.Collection
}

func (dbClient Adapter) FindUserByEmail(email *string) (*user_models.User, error) {
	var user user_models.User
	err := dbClient.usersCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		log.Printf("Error in finding user: %v", err)
		return &user_models.User{}, err
	}
	return &user, nil
}
func NewAdapter(uri string) *Adapter {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	fmt.Println("Connected Succesfully and pinged MongoDB!")
	db := client.Database("naresh-apps")
	usersCollection := db.Collection("users")
	computeCollection := db.Collection("compute")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	if _, err := usersCollection.Indexes().CreateOne(context.Background(), indexModel); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n Unique index on 'email' created succesfully.")
	return &Adapter{
		db:                db,
		usersCollection:   usersCollection,
		computeCollection: computeCollection,
	}
}
func (dbClient Adapter) CloseDBConnection() {
	// implements db close connection
	log.Output(1, "Closing the db connection")
}
func (dbClient Adapter) SaveCompute(params *types.CalcParameters, result int, operation string) (*models.Calculation, error) {
	calculation := &models.Calculation{
		X:         params.X,
		Y:         params.Y,
		Operation: operation,
		Result:    result,
	}
	res, err := dbClient.computeCollection.InsertOne(context.TODO(), calculation)
	fmt.Println(res)
	return calculation, err
}
