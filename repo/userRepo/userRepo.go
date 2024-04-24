package userRepo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"socialnetwork/domain"
	"socialnetwork/repo/dbconnections"
)

type UserRepo struct {
	*dbconnections.Db
}

func (r *UserRepo) FindById(id string) *domain.Usr {
	result := domain.Usr{}
	fter := bson.D{{"id", id}}
	err := r.MongoDB.UserCollection.FindOne(r.MongoDB.Context, fter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return &result

}

func (r *UserRepo) SaveUser(u *domain.Usr) *domain.Usr {
	id := u.Id
	ff := bson.D{{"id", id}}
	userFromDb := domain.Usr{}
	// Check if MongoDB collection is nil
	if r.MongoDB == nil {
		// Handle error
		fmt.Println("MongoDB connection is nil")
		return nil
	}

	if r.MongoDB.UserCollection == nil {
		fmt.Println("MongoDB UserCollection is nil")
		return nil
	}

	// Query MongoDB collection
	result := r.MongoDB.UserCollection.FindOne(r.MongoDB.Context, ff)
	if result != nil {
		// Decode result into userFromDb
		err := result.Decode(&userFromDb)
		if err != nil {
			// Handle decode error
			fmt.Println("Error decoding result:", err)
		}
		// Delete existing user
		r.DeleteOne(&userFromDb)

	}

	// Insert new user
	_, err := r.MongoDB.UserCollection.InsertOne(r.MongoDB.Context, u)
	if err != nil {
		// Handle insert error
		fmt.Println("Error inserting user:", err)
		return nil
	}

	return r.FindById(u.Id)
}

func (r *UserRepo) DeleteOne(u *domain.Usr) {
	usrBson, err := bson.Marshal(u)
	checkErr(err)
	r.MongoDB.UserCollection.DeleteOne(r.MongoDB.Context, usrBson)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
