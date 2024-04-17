package userRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fter := bson.M{"id": id}
	userFromDb := domain.Usr{}
	opt := options.FindOne()
	result := r.MongoDB.UserCollection.FindOne(context.TODO(), fter, opt)
	if result != nil {
		result.Decode(&userFromDb)
		r.DeleteONe(&userFromDb)
	}
	/*if err == nil {
		r.DeleteONe(&userFromDb)
	}*/
	r.MongoDB.UserCollection.InsertOne(r.MongoDB.Context, u)
	return r.FindById(u.Id)
}

func (r *UserRepo) DeleteONe(u *domain.Usr) {
	usrBson, err := bson.Marshal(u)
	checkErr(err)
	r.MongoDB.UserCollection.DeleteOne(r.MongoDB.Context, usrBson)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
