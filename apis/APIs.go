package main


import (
    "net/http"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gin-gonic/gin"
    "context"
	"log"
    "math"
    "os"
)

// data


type BdData struct {
    ID  primitive.ObjectID  `bson:"_id,omitempty"`
    X  float64  `json:"x"`
    Y float64  `json:"y"`
    Result float64 `json:"result"`
    Microservice  string `json:"microservice"`
}

type BdDataInsert struct {
    X  float64  `json:"x"`
    Y float64  `json:"y"`
    Result float64 `json:"result"`
    Microservice  string `json:"microservice"`
}

type RequestData struct {
    X  float64  `json:"x"`
    Y float64  `json:"y"`
}

var collection *mongo.Collection
var ctx = context.TODO()
var APIID = "A"

func main() {

    APIID = os.Getenv("process")

    router := gin.Default()
    router.GET("/getData/", getData)
    router.POST("/insert/", insert)

    initBD()

    router.Run("0.0.0.0:8080")
}

func initBD(){
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    clientOptions := options.Client().ApplyURI("mongodb+srv://luishenriques:luis123@trabsd.owhqaht.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI) //Server do mongo remoto para não baixar o bd
    client, err := mongo.Connect(ctx, clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

    collection = client.Database("Microservices").Collection("data")
}

// getAlbums responds with the list of all insertion in db as JSON.
func getData(c *gin.Context) {

    var AllBdData = []BdData{}

    cursor, err := collection.Find(context.TODO(), bson.D{})

    if err != nil {
        log.Fatal(err)
    }

    for cursor.Next(context.TODO()) {

        var result BdData
        if err := cursor.Decode(&result); err != nil {
            log.Fatal(err)
        }

        AllBdData = append(AllBdData,result)
    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    c.IndentedJSON(http.StatusOK, AllBdData)
}


// postAlbums adds an album from JSON received in the request body.
func insert(c *gin.Context) {
    // newData,err := io.ReadAll(c.Request.Body)
    var requestBody RequestData
    if err := c.BindJSON(&requestBody); err != nil {
        log.Fatal(err)
    }
    var teste = math.Pow(requestBody.X, requestBody.Y)


    insertData := BdDataInsert{X: requestBody.X, Y: requestBody.Y, Result: teste, Microservice: APIID}

    result, err := collection.InsertOne(context.Background(), insertData)

    if err != nil {
        log.Fatal(err)
    }

    filter := bson.D{{"_id", result.InsertedID}}
    var object BdData
    err = collection.FindOne(context.TODO(), filter).Decode(&object)

    c.IndentedJSON(http.StatusCreated, object)
}