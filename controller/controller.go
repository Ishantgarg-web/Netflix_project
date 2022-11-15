package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ishant-tata/NetFlix_Project_MongoDB/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection_string = "mongodb+srv://<username>:<password>@cluster0.vy1hm7p.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const collectionName = "watchlist"

// Reference of mongoDB collection
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection_string))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Successfully...")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("MongoDb instance is ready...")
}

/******     MongoDB helper function - CRUD    ***********/

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	res, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document is inserted Successfully with id: ", res.InsertedID)
}

// update 1 record (change false to true of watched) updateusing Id.
func updateOneMovie(movieId string) {
	// first need to change the give string into _id,
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"watched", true}}}}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document is updated successfully", res.MatchedCount)

}

// Delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.D{{"_id", id}}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Record successfully with count: ", res.DeletedCount)
}

// Delete Many Records
func deleteManyMovie() {
	filter := bson.D{{}}
	res, err := collection.DeleteMany(context.TODO(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All records Deleted successfully with count: ", res.DeletedCount)
}

// GetAllMovies
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var movies []primitive.M

	for cursor.Next(context.TODO()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

/********  methods that will use in routing ************/

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode("Movie is Created")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	s := "Movie with id " + params["id"] + " updated successfully"
	json.NewEncoder(w).Encode(s)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	s := "Movie with id " + params["id"] + " deleted successfully"
	json.NewEncoder(w).Encode(s)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deleteManyMovie()
	json.NewEncoder(w).Encode("All Movies Deleted Successfully")
}
