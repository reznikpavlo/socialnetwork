package messageRepo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"socialnetwork/domain"
	"socialnetwork/repo/dbconnections"
)

type MessageRepoMongo struct {
	*dbconnections.Db
}

func (m *MessageRepoMongo) FindById(id int64) domain.Message {
	result := domain.Message{}
	fter := bson.D{{"id", id}}
	err := m.MongoDB.MessageCollection.FindOne(m.MongoDB.Context, fter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
	}
	return result

}

func (m *MessageRepoMongo) FindAll() []domain.Message {
	return m.FindTop(100)
}

func (m *MessageRepoMongo) FindTop(top int64) []domain.Message {
	var resultMessages []domain.Message
	m.MongoDB.ConnectDB()
	findOptions := options.Find()
	if top != 0 {
		findOptions = options.Find().SetLimit(top)
	}
	cursor, err := m.MongoDB.MessageCollection.Find(m.MongoDB.Context, bson.D{}, findOptions)
	checkErr(err)
	defer cursor.Close(m.MongoDB.Context)

	if err = cursor.All(m.MongoDB.Context, &resultMessages); err != nil {
		panic(err)
	}

	/*	for _, result := range resultMessages {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}*/

	return resultMessages
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (m *MessageRepoMongo) Save(message *domain.Message) domain.Message {

	if message.Id == 0 {
		lastMessage := m.GetLastMessage()
		message.Id = lastMessage.Id + 1
		one, err := m.MongoDB.MessageCollection.InsertOne(m.MongoDB.Context, message)
		fmt.Println(one)
		fmt.Println(err)
	} else {
		messageBSON := bson.D{{"$set", bson.D{{"text", message.Text}}}}
		fter := bson.D{{"id", message.Id}}
		updateResult, err := m.MongoDB.MessageCollection.UpdateOne(m.MongoDB.Context, fter, messageBSON)
		log.Println("matched count = ", updateResult.MatchedCount)
		log.Println("modified count = ", updateResult.ModifiedCount)
		checkErr(err)
	}

	return m.FindById(message.Id)

}

func (m *MessageRepoMongo) DeleteOne(message *domain.Message) {
	messageBSON, err := bson.Marshal(message)
	checkErr(err)
	m.MongoDB.MessageCollection.DeleteOne(m.MongoDB.Context, messageBSON)
}

func (m *MessageRepoMongo) GetLastMessage() *domain.Message {
	fter := bson.D{}
	var messages []domain.Message
	optSort := options.Find().SetSort(bson.D{{"id", -1}})
	cur, _ := m.MongoDB.MessageCollection.Find(m.MongoDB.Context, fter, optSort)

	cur.All(m.MongoDB.Context, &messages)
	return &messages[0]

}
