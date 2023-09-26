package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vishnuavm5:pG0KuZfvAF711cwA@cluster0.nnnmhjk.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"

const colName = "watchlist"

//MOST IMPORTANT

var collection *mongo.Collection

//connect with mongodb

//only runs first time

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance

	fmt.Println("Collection instace is ready")

}

//MONGODB helpers

//insert one record

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Inserted one movie with db with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted successfully and the count is :", result.DeletedCount)
}

// delet all records
func deleteAllMovie() int64 {

	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", result.DeletedCount)

	return result.DeletedCount

}

//get all movies

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies

}

// actual controllers
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
