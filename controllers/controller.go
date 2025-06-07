package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	Todos "mongoapi/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var connectionStr string = "mongodb://localhost:27017/"
var dbName string = "todos"
var colName string = "todos"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionStr)

	client, _ := mongo.Connect(clientOptions)

	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to mongo db successfully")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}

func createTodo(todo Todos.Todo) {
	fmt.Println("Emntering 2")

	createdTodo, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created one todo in db", createdTodo)
}

func updateTodo(todoId string) {
	id, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"checked": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(result)
	}

	fmt.Println("Updated todo", result.MatchedCount)
}

func deleteOneTodo(todoId string) {
	id, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo deleted with delete count: ", deleteCount)
}

func deleteManyTodo() {
	deletedCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deletedCount)
}

func getAllTodos() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var todos []bson.M

	for cursor.Next(context.Background()) {
		var todo bson.M
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	defer cursor.Close(context.Background())
	return todos
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	allTodos := getAllTodos()
	json.NewEncoder(w).Encode(allTodos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var todo Todos.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}
	createTodo(todo)
	fmt.Println("Emntering")
	json.NewEncoder(w).Encode(todo)
}

func CheckTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateTodo(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneTodo(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteManyTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteManyTodo()
	json.NewEncoder(w).Encode("Deleted successfully")
}
