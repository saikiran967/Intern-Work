//Aim To perform crud operations on mongodb using go lang
package main
import (
"context" 
"fmt" 
"os" 
"reflect" 
"time"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/bson"
)
type student struct {
	 Name string 
	 Roll int 
    City string 
}
func main(){
	var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
	fmt.Println("mongo.Connect() ERROR:", err)
	os.Exit(1)
	} else{
		fmt.Println("Connected to mongoDB")
	}
	collection := client.Database("school").Collection("student")
	fmt.Println("collection type:", reflect.TypeOf(collection), "\n")
	a := student{"prakash", 1, "sap"}
	b := student{"siva", 2, "cpt"}
	c := student{"sai", 3, "gnt"}
 	s := []interface{}{a,b,c}
	a.create(client,collection,s)
	a.read(client,collection,s)
	a.update(client,collection,s)
	a.delete(client,collection,s)
	err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
	
}
func(ref student) create(client *mongo.Client,collection *mongo.Collection,s []interface{}){
	fmt.Println(ref.Name)
	_, e := collection.InsertMany(context.TODO(), s)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Documents inserted ")
	}
}
func(ref student) read(client *mongo.Client,collection *mongo.Collection,s []interface{}){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	var result student
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	} else{
	fmt.Println(result)
	}
	cursor, err := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
    var i bson.M
	if err = cursor.Decode(&i); err != nil {
        fmt.Println(err)
    }
    fmt.Println(i)
	}
}
func(ref student) update(client *mongo.Client,collection *mongo.Collection,s []interface{}){
	filter := bson.D{{"name", "siva"}}
 	update := bson.D{
    {"$set", bson.D{
        {"city", "sap"},
    }},
	}
 	_, e := collection.UpdateOne(context.TODO(), filter, update)
	var result student
	e = collection.FindOne(context.TODO(), bson.D{{"name","siva"}}).Decode(&result)
	fmt.Println(result)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Document updated")
	}
}

func(ref student) delete(client *mongo.Client,collection *mongo.Collection,s []interface{}){
	_, e := collection.DeleteMany(context.TODO(), bson.D{{}})
	
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Documents deleted")
	}
}