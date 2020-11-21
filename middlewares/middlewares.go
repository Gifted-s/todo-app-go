package middlewares



import (
    "fmt"
	 "time"
    "strconv"
	"net/http"
	"github.com/gorilla/mux"
	 "log"
	 "context"
	 "../models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)


 
func connectDb () *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb+srv://sunkanmi:sunkanmi123@cluster0.sq9uz.mongodb.net/todo-app?retryWrites=true&w=majority")

   client, err := mongo.Connect(context.TODO(), clientOptions)

   if err != nil {

	   log.Fatal("fail to connect to db")

   }
   fmt.Println("MongoDB connected")

   collection := client.Database("todo_api").Collection("todo_list")

   return collection


}

var collection = connectDb()

var year, month, day = time.Now().Date()

var todaysDate = month.String() + " " +  strconv.Itoa(day) + " " + strconv.Itoa(year)

func HandleHome(res http.ResponseWriter, req *http.Request){

res.Header().Set("Content-Type", "application/json");

json.NewEncoder(res).Encode("Hello world")
}
func HandleCreateList(res http.ResponseWriter, req *http.Request){

res.Header().Set("Content-Type", "application/json");

var list models.List;

json.NewDecoder(req.Body).Decode(&list)

list.DateCreated =  todaysDate

list.Tasks[0].DateCreated= todaysDate

list.Tasks[0].ID = primitive.NewObjectID()

result, err :=  collection.InsertOne(context.TODO(), list)

if err != nil {
  log.Fatal("insertion failed")
}

json.NewEncoder(res).Encode(result)
}

func HandleGetLists(res http.ResponseWriter, req *http.Request){

  res.Header().Set("Content-Type", "application/json");

  var lists []models.List;

  cursor, err := collection.Find(context.TODO(), bson.M{})

  if err != nil {
	  log.Fatal("failed to find all list")
  }

  defer cursor.Close(context.TODO())

  for cursor.Next(context.TODO()){
	var list models.List
	err := cursor.Decode(&list)
	if err != nil {
	  log.Fatal("failed to decode list")
  }
	lists= append(lists, list)
  }


  json.NewEncoder(res).Encode(lists)
 }


 func HandleEditList(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json");
	var list models.List;
	params := mux.Vars(req)
	_id, _:= primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(req.Body).Decode(&list)
	filter := bson.M{"_id": _id}
    update := bson.M{
		"$set": bson.M{"name":list.Name },
	}
	err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)
    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(list)
		
}
func HandleEditTask(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json");
	var list models.List;
	var task models.Task;
	params := mux.Vars(req)
	list_id, _:= primitive.ObjectIDFromHex(params["id"])
	task_id, _:= primitive.ObjectIDFromHex(params["task_id"])
	json.NewDecoder(req.Body).Decode(&task)
	filter := bson.M{"_id": list_id, "tasks._id": task_id}

	update := bson.M{
		
		"$set": bson.M{"tasks.$.name":task.Name },
	}
	err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)
    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(list)
		
}
func HandleCompleteTask(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json");
	var list models.List;
	var task models.Task;
	params := mux.Vars(req)
	list_id, _:= primitive.ObjectIDFromHex(params["id"])
	task_id, _:= primitive.ObjectIDFromHex(params["task_id"])
	json.NewDecoder(req.Body).Decode(&task)
	filter := bson.M{"_id": list_id, "tasks._id": task_id}

	update := bson.M{
		
		"$set": bson.M{"tasks.$.completed":true },
	}
	err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)
    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(list)
		
}

func HandleUndoTask(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json");
	var list models.List;
	var task models.Task;
	params := mux.Vars(req)
	list_id, _:= primitive.ObjectIDFromHex(params["id"])
	task_id, _:= primitive.ObjectIDFromHex(params["task_id"])
	json.NewDecoder(req.Body).Decode(&task)
	filter := bson.M{"_id": list_id, "tasks._id": task_id}

	update := bson.M{
		
		"$set": bson.M{"tasks.$.completed":false },
	}
	err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)
    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(list)
		
}

func HandleDeleteTask(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json");
	var list models.List
	params := mux.Vars(req)
    
	list_id, _:= primitive.ObjectIDFromHex(params["id"])

	task_id, _:= primitive.ObjectIDFromHex(params["task_id"])
	filter := bson.M{"_id": list_id}

	update := bson.M{"$pull": bson.M{"tasks": bson.M{"_id": task_id}},}
    err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)

    if err !=nil {
		log.Fatal(err)
	}
	
	json.NewEncoder(res).Encode(list)
		
}
func HandleAddTask(res http.ResponseWriter, req *http.Request){

	res.Header().Set("Content-Type", "application/json");

	var task models.Task;

	var list models.List;

	params := mux.Vars(req)
	_id, _:= primitive.ObjectIDFromHex(params["id"])

	json.NewDecoder(req.Body).Decode(&task)

	filter := bson.M{"_id": _id}

	task.ID = primitive.NewObjectID()
	task.DateCreated = todaysDate
	
    update := bson.M{

		"$push": bson.M{"tasks":task },
	}
	err:= collection.FindOneAndUpdate(context.TODO(),  filter, update).Decode(&list)

    if err !=nil {

		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(list)
		
}
func HandleDelete(res http.ResponseWriter, req *http.Request){

	res.Header().Set("Content-Type", "application/json");

	params := mux.Vars(req)

	_id, _:= primitive.ObjectIDFromHex(params["id"])
	
	
	filter := bson.M{"_id": _id}

	result, err:= collection.DeleteOne(context.TODO(),  filter)
    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}
	
	json.NewEncoder(res).Encode(result)		
}
func HandleGetList(res http.ResponseWriter, req *http.Request){

	var list models.List

	res.Header().Set("Content-Type", "application/json");

	params := mux.Vars(req)

	_id, _:= primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": _id}

	err:= collection.FindOne(context.TODO(),  filter).Decode(&list)

    if err !=nil {
		log.Fatal("Update Operation not successfull")
	}

	json.NewEncoder(res).Encode(list)		
}


   
